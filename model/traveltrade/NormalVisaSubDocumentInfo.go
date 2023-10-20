package traveltrade

import (
	"sync"
)

// NormalVisaSubDocumentInfo 结构体
type NormalVisaSubDocumentInfo struct {
	// 文件类型
	FileType string `json:"file_type,omitempty" xml:"file_type,omitempty"`
	// base64编码的文件内容
	FileContent string `json:"file_content,omitempty" xml:"file_content,omitempty"`
	// 子材料文档编号，10001:护照封面，10002:护照首页，11401:去程机票，11402:返程机票
	DocType int64 `json:"doc_type,omitempty" xml:"doc_type,omitempty"`
}

var poolNormalVisaSubDocumentInfo = sync.Pool{
	New: func() any {
		return new(NormalVisaSubDocumentInfo)
	},
}

// GetNormalVisaSubDocumentInfo() 从对象池中获取NormalVisaSubDocumentInfo
func GetNormalVisaSubDocumentInfo() *NormalVisaSubDocumentInfo {
	return poolNormalVisaSubDocumentInfo.Get().(*NormalVisaSubDocumentInfo)
}

// ReleaseNormalVisaSubDocumentInfo 释放NormalVisaSubDocumentInfo
func ReleaseNormalVisaSubDocumentInfo(v *NormalVisaSubDocumentInfo) {
	v.FileType = ""
	v.FileContent = ""
	v.DocType = 0
	poolNormalVisaSubDocumentInfo.Put(v)
}
