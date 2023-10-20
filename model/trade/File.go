package trade

import (
	"sync"
)

// File 结构体
type File struct {
	// 返回的是绝对路径如：http://img07.taobaocdn.com/imgextra/i7/22670458/T2dD0kXb4cXXXXXXXX_!!22670458.jpg
	FilePath string `json:"file_path,omitempty" xml:"file_path,omitempty"`
	// 图片状态,unfroze代表没有被冻结，froze代表被冻结,pass代表排查通过
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 文件是否删除
	Deleted string `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 图片像素，如果非图片，该值为空
	PicturePix string `json:"picture_pix,omitempty" xml:"picture_pix,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 文件的大小
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}

var poolFile = sync.Pool{
	New: func() any {
		return new(File)
	},
}

// GetFile() 从对象池中获取File
func GetFile() *File {
	return poolFile.Get().(*File)
}

// ReleaseFile 释放File
func ReleaseFile(v *File) {
	v.FilePath = ""
	v.Status = ""
	v.Deleted = ""
	v.PicturePix = ""
	v.GmtCreate = ""
	v.Size = 0
	poolFile.Put(v)
}
