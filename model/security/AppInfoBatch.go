package security

import (
	"sync"
)

// AppInfoBatch 结构体
type AppInfoBatch struct {
	// 需要扫描的应用的具体信息列表
	ScanInfos []AppInfoBatchItem `json:"scan_infos,omitempty" xml:"scan_infos>app_info_batch_item,omitempty"`
	// 回调地址,dataType=4时必填,用于处理完成后反向通知,通知为GET请求,请求格式:  callbackUrl+&#34;?itemId=xxx&amp;taskStatus=1&#34;
	CallbackUrl string `json:"callback_url,omitempty" xml:"callback_url,omitempty"`
	// APP应用类型 1-android 2-ios(暂不支持)
	AppOsType int64 `json:"app_os_type,omitempty" xml:"app_os_type,omitempty"`
	// APP数据类型 3-Batch MD5 4-Batch URL(暂不支持)
	DataType int64 `json:"data_type,omitempty" xml:"data_type,omitempty"`
}

var poolAppInfoBatch = sync.Pool{
	New: func() any {
		return new(AppInfoBatch)
	},
}

// GetAppInfoBatch() 从对象池中获取AppInfoBatch
func GetAppInfoBatch() *AppInfoBatch {
	return poolAppInfoBatch.Get().(*AppInfoBatch)
}

// ReleaseAppInfoBatch 释放AppInfoBatch
func ReleaseAppInfoBatch(v *AppInfoBatch) {
	v.ScanInfos = v.ScanInfos[:0]
	v.CallbackUrl = ""
	v.AppOsType = 0
	v.DataType = 0
	poolAppInfoBatch.Put(v)
}
