package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	// หมายเหตุ: ตั้งแต่ Go 1.20 เป็นต้นไป การเรียก rand.Seed ไม่จำเป็นแล้ว
	// เพราะ Go จะสุ่ม Seed ให้อัตโนมัติ แต่ใส่ไว้เผื่อคุณใช้ Go เวอร์ชันเก่าครับ
	rand.Seed(time.Now().UnixNano())
}

// RandomInt สุ่มตัวเลขระหว่าง min ถึง max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString สุ่มตัวอักษรความยาว n ตัว
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// -------------------------------------------------------------------
// ฟังก์ชันสุ่มสำหรับแต่ละ Field
// -------------------------------------------------------------------

// RandomUsername สุ่ม Username (ความยาว 8 ตัวอักษร)
func RandomUsername() string {
	return RandomString(8)
}

// RandomEmail สุ่ม Email แอดเดรส
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}

// RandomPassword สุ่มรหัสผ่าน (ความยาว 10 ตัวอักษร)
func RandomPassword() string {
	return RandomString(10)
}

// RandomDisplayName สุ่มชื่อที่ใช้แสดงผล
func RandomDisplayName() string {
	return fmt.Sprintf("User_%s", RandomString(5))
}

// RandomBio สุ่มประวัติย่อ (Bio)
func RandomBio() string {
	bios := []string{
		"Hello world!",
		"I love coding in Go.",
		"Just a random user.",
		"Handsome and smart.",
		"Learning backend development.",
	}
	return bios[rand.Intn(len(bios))]
}

// RandomAvatarURL สุ่มลิงก์รูปภาพโปรไฟล์
func RandomAvatarURL() string {
	return fmt.Sprintf("https://example.com/avatar/%s.jpg", RandomString(10))
}
