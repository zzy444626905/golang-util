package fileutil

import "os"

// IsDir Determine whether the path is a folder, true: folder; false: file;
func IsDir(s string) bool {
	open, err := os.Open(s)
	if err != nil {
		return false
	}
	stat, err := open.Stat()
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// IsFile Determine whether the path is a file, true: file; false: folder;
func IsFile(s string) bool {
	return !IsDir(s)
}
