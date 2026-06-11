package db

import (
	"context"
	"testing"
	"time"

	"github.com/DefinitelyNotJay/flyte/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T, user User) Post {
	arg := CreatePostParams{
		AuthorID: user.ID,
		Content: pgtype.Text{
			String: util.RandomString(100),
			Valid:  true,
		},
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.AuthorID, post.AuthorID)
	require.Equal(t, arg.Content, post.Content)

	require.NotZero(t, post.ID)
	require.NotZero(t, post.CreatedAt)

	return post
}

func TestCreatePost(t *testing.T) {
	user := createRandomUser(t)
	createRandomPost(t, user)
}

func TestGetPost(t *testing.T) {
	user := createRandomUser(t)
	post1 := createRandomPost(t, user)
	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.AuthorID, post2.AuthorID)
	require.Equal(t, post1.Content, post2.Content)
	require.WithinDuration(t, post1.CreatedAt.Time, post2.CreatedAt.Time, time.Second)
}
