package storage

import (
	"os"
	"path/filepath"
	"time"
)

const baseDir = ".local/share/micro/posts"

func getStorageDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, baseDir)
}

func SaveEncryptedPost(data []byte) (string, error) {
	dir := getStorageDir()
	os.MkdirAll(dir, 0700)

	filename := time.Now().Format("2006-01-02T15-04-05") + ".json.enc"
	path := filepath.Join(dir, filename)

	err := os.WriteFile(path, data, 0600)
	return filepath.Base(path), err
}

func LoadEncryptedPosts() ([]string, error) {
	dir := getStorageDir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}

func ReadEncryptedPost(filename string) ([]byte, error) {
	dir := getStorageDir()
	return os.ReadFile(filepath.Join(dir, filename))
}

func DeletePostFile(filename string) error {
	dir := getStorageDir()
	return os.Remove(filepath.Join(dir, filename))
}
