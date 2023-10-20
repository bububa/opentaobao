package security

import (
	"sync"
)

// UrlScanResult 结构体
type UrlScanResult struct {
	// 扫描详细结果
	UrlResultItemList []UrlScanResultItem `json:"url_result_item_list,omitempty" xml:"url_result_item_list>url_scan_result_item,omitempty"`
	// 请求标志唯一id
	EventId string `json:"event_id,omitempty" xml:"event_id,omitempty"`
}

var poolUrlScanResult = sync.Pool{
	New: func() any {
		return new(UrlScanResult)
	},
}

// GetUrlScanResult() 从对象池中获取UrlScanResult
func GetUrlScanResult() *UrlScanResult {
	return poolUrlScanResult.Get().(*UrlScanResult)
}

// ReleaseUrlScanResult 释放UrlScanResult
func ReleaseUrlScanResult(v *UrlScanResult) {
	v.UrlResultItemList = v.UrlResultItemList[:0]
	v.EventId = ""
	poolUrlScanResult.Put(v)
}
