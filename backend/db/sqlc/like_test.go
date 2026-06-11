package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomLike(t *testing.T, user User, post Post) Like {
	arg := CreateLikeParams{
		UserID: user.ID,
		PostID: post.ID,
	}

	like, err := testQueries.CreateLike(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, like)

	require.Equal(t, arg.UserID, like.UserID)
	require.Equal(t, arg.PostID, like.PostID)
	require.NotZero(t, like.CreatedAt)

	return like
}

func TestCreateLike(t *testing.T) {
	user := createRandomUser(t)
	post := createRandomPost(t, user)
	createRandomLike(t, user, post)
}

func TestGetLike(t *testing.T) {
	user := createRandomUser(t)
	post := createRandomPost(t, user)
	like1 := createRandomLike(t, user, post)

	like2, err := testQueries.GetLike(context.Background(), GetLikeParams{
		UserID: user.ID,
		PostID: post.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, like2)

	require.Equal(t, like1.UserID, like2.UserID)
	require.Equal(t, like1.PostID, like2.PostID)
	require.WithinDuration(t, like1.CreatedAt.Time, like2.CreatedAt.Time, time.Second)
}

func TestDeleteLike(t *testing.T) {
	user := createRandomUser(t)
	post := createRandomPost(t, user)
	createRandomLike(t, user, post)

	err := testQueries.DeleteLike(context.Background(), DeleteLikeParams{
		UserID: user.ID,
		PostID: post.ID,
	})
	require.NoError(t, err)

	like2, err := testQueries.GetLike(context.Background(), GetLikeParams{
		UserID: user.ID,
		PostID: post.ID,
	})
	require.Error(t, err)
	require.Empty(t, like2)
}

func TestListLikesByPost(t *testing.T) {
	user := createRandomUser(t)
	post := createRandomPost(t, user)

	for i := 0; i < 5; i++ {
		u := createRandomUser(t)
		createRandomLike(t, u, post)
	}

	arg := ListLikesByPostParams{
		PostID: post.ID,
		Limit:  5,
		Offset: 0,
	}

	likes, err := testQueries.ListLikesByPost(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, likes, 5)

	for _, like := range likes {
		require.NotEmpty(t, like)
		require.Equal(t, post.ID, like.PostID)
	}
}
