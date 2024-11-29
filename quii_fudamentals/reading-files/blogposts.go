package blogposts

import (
	"io/fs"
	"testing/fstest"
)

type Post struct {
	Title string
}

func NewPostsFromFS(fileSystem fstest.MapFS) ([]Post, error) {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts, nil
}
