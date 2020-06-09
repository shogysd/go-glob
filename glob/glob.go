// package overview
// Recursively scans the files under the specified directory and returns them as slice pointers.
// Ignored if the hidden file (first letter is '.')
// arguments: (path string, filePaths *[]string, fileNamePatternReg string, filePathPatternReg string)
package glob

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Glob(path string, filePaths *[]string, fileNamePatternReg string, filePathPatternReg string) (retErr error) {
	var err error
	var inDirFilePaths []string

	if path[len(path)-1:] != "/" {
		path += "/"
	}

	if inDirFilePaths, err = filepath.Glob(path + "*"); err != nil {
		fmt.Println("error")
	}

	for _, filePath := range inDirFilePaths {
		filename := strings.Split(filePath, "/")[len(strings.Split(filePath, "/"))-1]
		if filename[:1] != "." {
			fInfo, _ := os.Stat(filePath)
			if fInfo.IsDir() {
				Glob(filePath, filePaths, fileNamePatternReg, filePathPatternReg)
			} else if regexp.MustCompile(fileNamePatternReg).MatchString(filename) && regexp.MustCompile(filePathPatternReg).MatchString(filePath) {
				*filePaths = append(*filePaths, filePath)
			}
		}
	}
	return nil
}
