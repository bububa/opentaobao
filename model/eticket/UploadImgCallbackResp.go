package eticket

import (
	"sync"
)

// UploadImgCallbackResp 结构体
type UploadImgCallbackResp struct {
	// 扩展属性
	AttributeMap string `json:"attribute_map,omitempty" xml:"attribute_map,omitempty"`
	// 图片在淘宝的文件名
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
}

var poolUploadImgCallbackResp = sync.Pool{
	New: func() any {
		return new(UploadImgCallbackResp)
	},
}

// GetUploadImgCallbackResp() 从对象池中获取UploadImgCallbackResp
func GetUploadImgCallbackResp() *UploadImgCallbackResp {
	return poolUploadImgCallbackResp.Get().(*UploadImgCallbackResp)
}

// ReleaseUploadImgCallbackResp 释放UploadImgCallbackResp
func ReleaseUploadImgCallbackResp(v *UploadImgCallbackResp) {
	v.AttributeMap = ""
	v.FileName = ""
	poolUploadImgCallbackResp.Put(v)
}
