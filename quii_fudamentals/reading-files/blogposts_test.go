package blogposts_test

import (
	blogposts "https://github.com/andreivisan/learn_go/tree/master/quii_fudamentals/reading-files"
	"testing"
	"testing/fstest"
)

func TestBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md": {Data: []byte("Hi")},
		"hello-world.md": {Data: []byte("Hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
