package security

import (
	"sync"
)

// UrlScanParamItem 结构体
type UrlScanParamItem struct {
	// 需要扫描的url
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 参数标记，用于识别返回结果对应的参数
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
}

var poolUrlScanParamItem = sync.Pool{
	New: func() any {
		return new(UrlScanParamItem)
	},
}

// GetUrlScanParamItem() 从对象池中获取UrlScanParamItem
func GetUrlScanParamItem() *UrlScanParamItem {
	return poolUrlScanParamItem.Get().(*UrlScanParamItem)
}

// ReleaseUrlScanParamItem 释放UrlScanParamItem
func ReleaseUrlScanParamItem(v *UrlScanParamItem) {
	v.Data = ""
	v.Flag = ""
	poolUrlScanParamItem.Put(v)
}
