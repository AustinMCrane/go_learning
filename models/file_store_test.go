package models

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ret := m.Run()
	os.RemoveAll(RootFilePath + "/*")
	os.Exit(ret)
}
func TestSaveString(t *testing.T) {
	s := "Hello, World"
	path, err := SaveString(s)
	if err != nil {
		t.Error("Writer error")
	}
	if path != RootFilePath+"/something.txt" {
		t.Error("not correct file path")
	}
}

func TestSaveBase64String(t *testing.T) {
	img := "iVBORw0KGgoAAAANSUhEUgAAAAoAAAALCAYAAABGbhwYAAAAAXNSR0IArs4c6QAAAAlwSFlzAAALEwAACxMBAJqcGAAAAVlpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDUuNC4wIj4KICAgPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICAgICAgPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIKICAgICAgICAgICAgeG1sbnM6dGlmZj0iaHR0cDovL25zLmFkb2JlLmNvbS90aWZmLzEuMC8iPgogICAgICAgICA8dGlmZjpPcmllbnRhdGlvbj4xPC90aWZmOk9yaWVudGF0aW9uPgogICAgICA8L3JkZjpEZXNjcmlwdGlvbj4KICAgPC9yZGY6UkRGPgo8L3g6eG1wbWV0YT4KTMInWQAAAKBJREFUGBl1jssJQjEQRROxCkHEMgQbcGERaokubECwjacbbcEP6jkhyXtZvAsnuTO5wySGXhE76cvkvpy/Yc/QmNKbh5SpFX6bJ47cl+ybTTuaH3BI3rCHRnOqGxh4wTP7K/cC6ueX+Bn4edeIQ/Z8q8EO/8i16yXCHTpodKAa/lFvL8kp17hSrWGTXAgn7nP2ZpIcGFMNlYDhaSmyr6E/JAAfgeFNqDQAAAAASUVORK5CYII="

	path, err := SaveBase64Image("test.png", img)
	if err != nil {
		t.Error("Writer error")
	}
	if path != RootFilePath+"/test.png" {
		t.Error("not correct file path")
	}
}
