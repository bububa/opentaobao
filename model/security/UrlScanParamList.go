package security

import (
	"sync"
)

// UrlScanParamList 结构体
type UrlScanParamList struct {
	// 扫描参数列表
	UrlScanParamItems []UrlScanParamItem `json:"url_scan_param_items,omitempty" xml:"url_scan_param_items>url_scan_param_item,omitempty"`
}

var poolUrlScanParamList = sync.Pool{
	New: func() any {
		return new(UrlScanParamList)
	},
}

// GetUrlScanParamList() 从对象池中获取UrlScanParamList
func GetUrlScanParamList() *UrlScanParamList {
	return poolUrlScanParamList.Get().(*UrlScanParamList)
}

// ReleaseUrlScanParamList 释放UrlScanParamList
func ReleaseUrlScanParamList(v *UrlScanParamList) {
	v.UrlScanParamItems = v.UrlScanParamItems[:0]
	poolUrlScanParamList.Put(v)
}
