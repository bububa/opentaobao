package foodscan

import (
	"sync"
)

// FilePackageBasicReq 结构体
type FilePackageBasicReq struct {
	// 脚型报告的唯一标识
	ScanId string `json:"scan_id,omitempty" xml:"scan_id,omitempty"`
	// 第一张图片的文件名
	FileName1 string `json:"file_name1,omitempty" xml:"file_name1,omitempty"`
	// 第二张图片的文件名
	FileName2 string `json:"file_name2,omitempty" xml:"file_name2,omitempty"`
	// 用户唯一标识，可以是淘宝用户ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 第三张图片的文件名
	FileName3 string `json:"file_name3,omitempty" xml:"file_name3,omitempty"`
	// 1左脚 2右脚
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolFilePackageBasicReq = sync.Pool{
	New: func() any {
		return new(FilePackageBasicReq)
	},
}

// GetFilePackageBasicReq() 从对象池中获取FilePackageBasicReq
func GetFilePackageBasicReq() *FilePackageBasicReq {
	return poolFilePackageBasicReq.Get().(*FilePackageBasicReq)
}

// ReleaseFilePackageBasicReq 释放FilePackageBasicReq
func ReleaseFilePackageBasicReq(v *FilePackageBasicReq) {
	v.ScanId = ""
	v.FileName1 = ""
	v.FileName2 = ""
	v.UserId = ""
	v.FileName3 = ""
	v.Type = 0
	poolFilePackageBasicReq.Put(v)
}
