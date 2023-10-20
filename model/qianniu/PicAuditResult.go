package qianniu

import (
	"sync"
)

// PicAuditResult 结构体
type PicAuditResult struct {
	// 经审核的图片链接列表
	AuditedUrls []string `json:"audited_urls,omitempty" xml:"audited_urls>string,omitempty"`
	// 不合法的图片链接序号
	IllegalIndex []string `json:"illegal_index,omitempty" xml:"illegal_index>string,omitempty"`
}

var poolPicAuditResult = sync.Pool{
	New: func() any {
		return new(PicAuditResult)
	},
}

// GetPicAuditResult() 从对象池中获取PicAuditResult
func GetPicAuditResult() *PicAuditResult {
	return poolPicAuditResult.Get().(*PicAuditResult)
}

// ReleasePicAuditResult 释放PicAuditResult
func ReleasePicAuditResult(v *PicAuditResult) {
	v.AuditedUrls = v.AuditedUrls[:0]
	v.IllegalIndex = v.IllegalIndex[:0]
	poolPicAuditResult.Put(v)
}
