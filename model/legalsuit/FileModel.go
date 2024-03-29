package legalsuit

import (
	"sync"
)

// FileModel 结构体
type FileModel struct {
	// 附件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 附件Key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 沟通时间
	SubmitDate string `json:"submit_date,omitempty" xml:"submit_date,omitempty"`
	// 沟通结果
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 沟通内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 沟通方
	Intro string `json:"intro,omitempty" xml:"intro,omitempty"`
	// 上传时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
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
	v.SubmitDate = ""
	v.Description = ""
	v.Content = ""
	v.Intro = ""
	v.CreateTime = ""
	poolFileModel.Put(v)
}
