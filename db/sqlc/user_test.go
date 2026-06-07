package db

import (
	"context"
	"testing"
	"time"

	"github.com/DefinitelyNotJay/flyte/util"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomPassword())
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:     util.RandomUsername(),
		Email:        util.RandomEmail(),
		PasswordHash: hashedPassword,
		DisplayName: pgtype.Text{
			String: util.RandomDisplayName(),
			Valid:  true,
		},
		Bio: pgtype.Text{
			String: util.RandomBio(),
			Valid:  true,
		},
		AvatarUrl: pgtype.Text{
			String: util.RandomAvatarURL(),
			Valid:  true,
		},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)
	require.Equal(t, arg.DisplayName, user.DisplayName)
	require.Equal(t, arg.Bio, user.Bio)
	require.Equal(t, arg.AvatarUrl, user.AvatarUrl)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	return user

}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	// 3. ตรวจสอบว่าต้องไม่มี Error และข้อมูลต้องไม่ว่างเปล่า
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	// 4. ตรวจสอบว่าข้อมูลที่ดึงมา (user2) ตรงกับข้อมูลตอนสร้าง (user1) ทุกประการ
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.Equal(t, user1.DisplayName, user2.DisplayName)
	require.Equal(t, user1.Bio, user2.Bio)
	require.Equal(t, user1.AvatarUrl, user2.AvatarUrl)

	// หมายเหตุ: สำหรับเวลา (Time) บางครั้ง Database กับ Go อาจจะมีความคลาดเคลื่อนระดับเสี้ยววินาที
	// การใช้ require.WithinDuration จะปลอดภัยกว่า require.Equal ครับ
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
	require.WithinDuration(t, user1.UpdatedAt.Time, user2.UpdatedAt.Time, time.Second)
}

func TestUpdateUser(t *testing.T) {
	// 1. สร้าง User จำลองขึ้นมาก่อน
	user1 := createRandomUser(t)

	// 2. เตรียมข้อมูลใหม่ที่ต้องการอัปเดต (สมมติว่าเปลี่ยนแค่ Username กับ Bio)
	arg := UpdateUserParams{
		ID:           user1.ID,              // ต้องส่ง ID ไปด้วยเพื่อให้รู้ว่าอัปเดตใคร
		Username:     util.RandomUsername(), // สุ่มชื่อใหม่
		Email:        user1.Email,           // ค่าเดิม
		PasswordHash: user1.PasswordHash,    // ค่าเดิม
		DisplayName:  user1.DisplayName,     // ค่าเดิม
		Bio: pgtype.Text{
			String: "I am updated!", // ข้อความใหม่
			Valid:  true,
		},
		AvatarUrl: user1.AvatarUrl, // ค่าเดิม
	}

	// 3. สั่งอัปเดตข้อมูล
	user2, err := testQueries.UpdateUser(context.Background(), arg)

	// 4. ตรวจสอบผลลัพธ์
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	// 5. เช็คว่าฟิลด์ที่เรา "เปลี่ยน" ถูกอัปเดตจริงๆ
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Bio, user2.Bio)

	// 6. เช็คว่าฟิลด์ที่เรา "ไม่ได้เปลี่ยน" ยังคงมีค่าเดิม
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.Equal(t, user1.DisplayName, user2.DisplayName)
	require.Equal(t, user1.AvatarUrl, user2.AvatarUrl)

	// 7. CreatedAt ต้องเป็นเวลาเดิม แต่ UpdatedAt อาจจะถูกเปลี่ยนโดย Database (ถ้าคุณตั้ง Trigger ไว้)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}
