package legalcase

import (
	"sync"
)

// FileModel 结构体
type FileModel struct {
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}

var poolFileModel = sync.Pool{
	New: func() any {
		return new(FileModel)
	},
}

// GetFileModel() 从对象池中获取FileModel
func GetFileModel() *FileModel {
	return poolFileModel.Get().(*FileModel)
}

// ReleaseFileModel 释放FileModel
func ReleaseFileModel(v *FileModel) {
	v.Name = ""
	v.Key = ""
	poolFileModel.Put(v)
}
