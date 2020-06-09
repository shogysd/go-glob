package test

import (
	"../glob"
	"fmt"
	"reflect"
	"testing"
)

func TestGlob(t *testing.T) {
	var filePaths []string

	// do NOT use regex
	var basicTestPassList = []string{"glob_test_source/dir1/file1-1",
		"glob_test_source/dir1/file1-2"}
	var multipleLayersTestPassList = []string{"glob_test_source/dir2/dir2-2/dir2-3/file2-3-1",
		"glob_test_source/dir2/dir2-2/dir2-3/file2-3-2",
		"glob_test_source/dir2/dir2-2/dir2-3/file2-3-3",
		"glob_test_source/dir2/dir2-2/file2-1-1",
		"glob_test_source/dir2/file2-1"}

	if err := glob.Glob("./glob_test_source/dir1", &filePaths, ".*", ".*"); err != nil {
		t.Error("glob run err (basic test)")
	} else if !reflect.DeepEqual(filePaths, basicTestPassList) {
		fmt.Println(filePaths)
		t.Error("error: basic test (last letter is NOT '/')")
	}
	filePaths = []string{}

	if err := glob.Glob("./glob_test_source/dir1/", &filePaths, ".*", ".*"); err != nil {
		t.Error("glob run err (basic test)")
	} else if !reflect.DeepEqual(filePaths, basicTestPassList) {
		fmt.Println(filePaths)
		t.Error("error: basic test (last letter is '/')")
	}
	filePaths = []string{}

	if err := glob.Glob("./glob_test_source/dir2", &filePaths, ".*", ".*"); err != nil {
		t.Error("glob run err (multiple layers test)")
	} else if !reflect.DeepEqual(filePaths, multipleLayersTestPassList) {
		fmt.Println(filePaths)
		t.Error("error: multiple layers test")
	}
	filePaths = []string{}

	if err := glob.Glob("./glob_test_source/dir3", &filePaths, ".*", ".*"); err != nil {
		t.Error("glob run err (hidden file and dir test)")
	} else if len(filePaths) != 0 {
		fmt.Println(filePaths)
		t.Error("error: hidden file and dir test")
	}
	filePaths = []string{}

	// use regex for filename
	var filenameRegTestPassList1 = []string{"glob_test_source/dir1/file1-1", "glob_test_source/dir1/file1-2"}
	var filenameRegTestPassList2 = []string{"glob_test_source/dir2/dir2-2/dir2-3/file2-3-3"}

	if err := glob.Glob("./glob_test_source", &filePaths, ".*le1-.*", ".*"); err != nil {
		t.Error("glob run err (hidden file and dir test)")
	} else if !reflect.DeepEqual(filePaths, filenameRegTestPassList1) {
		fmt.Println(filePaths)
		t.Error("error: hidden file and dir test")
	}
	filePaths = []string{}

	if err := glob.Glob("./glob_test_source", &filePaths, "^fi.*-.*3$", ".*"); err != nil {
		t.Error("glob run err (hidden file and dir test)")
	} else if !reflect.DeepEqual(filePaths, filenameRegTestPassList2) {
		fmt.Println(filePaths)
		t.Error("error: hidden file and dir test")
	}
	filePaths = []string{}

	// use regex for filepath
	var pathRegTestPassList = []string{"glob_test_source/dir2/dir2-2/dir2-3/file2-3-1",
		"glob_test_source/dir2/dir2-2/dir2-3/file2-3-2",
		"glob_test_source/dir2/dir2-2/dir2-3/file2-3-3",
		"glob_test_source/dir2/dir2-2/file2-1-1",
		"glob_test_source/dir2/file2-1"}
	if err := glob.Glob("./glob_test_source", &filePaths, ".*", ".*dir.*2-.*"); err != nil {
		t.Error("glob run err (hidden file and dir test)")
	} else if !reflect.DeepEqual(filePaths, pathRegTestPassList) {
		fmt.Println(filePaths)
		t.Error("error: hidden file and dir test")
	}
	filePaths = []string{}

	// use regex for filename and filepath
	var filenameAndPathRegTestPassList = []string{"glob_test_source/dir2/dir2-2/file2-1-1"}
	if err := glob.Glob("./glob_test_source", &filePaths, ".*-1-.*", ".*dir.*2-.*"); err != nil {
		t.Error("glob run err (hidden file and dir test)")
	} else if !reflect.DeepEqual(filePaths, filenameAndPathRegTestPassList) {
		fmt.Println(filePaths)
		t.Error("error: hidden file and dir test")
	}
	filePaths = []string{}
}
