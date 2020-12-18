package goldtest

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"syscall"
)

func removeSpecialCharacters(path string) string {
	path = strings.ReplaceAll(path, " ", "_")

	re := regexp.MustCompile("[^A-Za-z0-9/_-]")
	cleanedPath := re.ReplaceAllString(path, "")

	return cleanedPath
}

func preparePath(path string) string {
	return path + ".golden"
}

func prepareDirPath(path string) string {
	dir := strings.Split(path, "/")
	return strings.Join(dir[:len(dir)-1], "/")
}

func readFile(fileName string) ([]byte, error) {
	fileName = removeSpecialCharacters(fileName)

	expected, err := ioutil.ReadFile(preparePath(fileName))
	if err != nil {
		return nil, err
	}
	return expected, nil
}

func writeFile(fileName string, bytes []byte) error {
	syscall.Umask(0)
	fileName = removeSpecialCharacters(fileName)
	if dirPath := prepareDirPath(fileName); dirPath != "" {
		if _, err := os.Stat(dirPath); err != nil {
			if err := os.MkdirAll(dirPath, 0777); err != nil {
				return fmt.Errorf("Directory %s could not be created: %s", dirPath, err)
			}
		}
	}

	if err := ioutil.WriteFile(preparePath(fileName), bytes, 0666); err != nil {
		return fmt.Errorf("Error writing golden file for filename=%s: %s", preparePath(fileName), err)
	}
	return nil
}
