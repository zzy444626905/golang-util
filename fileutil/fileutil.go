package fileutil

import "os"

// IsDir Determine whether the path is a folder, true: folder; false: file;
func IsDir(s string) bool {
	if !Exists(s)  {
		return false
	}
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
	if !Exists(s)  {
		return false
	}
	return !IsDir(s)
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
