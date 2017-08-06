package models

import (
	"testing"
)

func TestCreatePost(t *testing.T) {
	body := "This is the body"
	img := "iVBORw0KGgoAAAANSUhEUgAAAAoAAAALCAYAAABGbhwYAAAAAXNSR0IArs4c6QAAAAlwSFlzAAALEwAACxMBAJqcGAAAAVlpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDUuNC4wIj4KICAgPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICAgICAgPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIKICAgICAgICAgICAgeG1sbnM6dGlmZj0iaHR0cDovL25zLmFkb2JlLmNvbS90aWZmLzEuMC8iPgogICAgICAgICA8dGlmZjpPcmllbnRhdGlvbj4xPC90aWZmOk9yaWVudGF0aW9uPgogICAgICA8L3JkZjpEZXNjcmlwdGlvbj4KICAgPC9yZGY6UkRGPgo8L3g6eG1wbWV0YT4KTMInWQAAAKBJREFUGBl1jssJQjEQRROxCkHEMgQbcGERaokubECwjacbbcEP6jkhyXtZvAsnuTO5wySGXhE76cvkvpy/Yc/QmNKbh5SpFX6bJ47cl+ybTTuaH3BI3rCHRnOqGxh4wTP7K/cC6ueX+Bn4edeIQ/Z8q8EO/8i16yXCHTpodKAa/lFvL8kp17hSrWGTXAgn7nP2ZpIcGFMNlYDhaSmyr6E/JAAfgeFNqDQAAAAASUVORK5CYII="
	post, err := CreatePost(body, img, "test.png")
	if err != nil {
		t.Error(err)
	}
	if post.Body != body {
		t.Error("body is not equal")
	}
}
