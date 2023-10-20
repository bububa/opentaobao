package simba

import (
	"sync"
)

// TaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult 结构体
type TaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopReportDownloadVO *TopReportDownloadVo `json:"top_report_download_v_o,omitempty" xml:"top_report_download_v_o,omitempty"`
}

var poolTaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult)
	},
}

// GetTaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult() 从对象池中获取TaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult
func GetTaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult() *TaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult {
	return poolTaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult.Get().(*TaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult)
}

// ReleaseTaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult 释放TaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult
func ReleaseTaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult(v *TaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult) {
	v.Info = nil
	v.TopReportDownloadVO = nil
	poolTaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult.Put(v)
}
