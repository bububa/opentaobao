package qianniu

import (
	"sync"
)

// PicAuditParam 结构体
type PicAuditParam struct {
	// 图片链接列表
	PicUrls []string `json:"pic_urls,omitempty" xml:"pic_urls>string,omitempty"`
	// 用户id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolPicAuditParam = sync.Pool{
	New: func() any {
		return new(PicAuditParam)
	},
}

// GetPicAuditParam() 从对象池中获取PicAuditParam
func GetPicAuditParam() *PicAuditParam {
	return poolPicAuditParam.Get().(*PicAuditParam)
}

// ReleasePicAuditParam 释放PicAuditParam
func ReleasePicAuditParam(v *PicAuditParam) {
	v.PicUrls = v.PicUrls[:0]
	v.UserId = 0
	poolPicAuditParam.Put(v)
}
