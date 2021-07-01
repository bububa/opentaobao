package idleparttime

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleParttimeSynclogAPIResponse
兼职同步日志 API返回值
alibaba.idle.parttime.synclog

提供给供应商查询的接口 */
type AlibabaIdleParttimeSynclogAPIResponse struct {
	model.CommonResponse
	AlibabaIdleParttimeSynclogAPIResponseModel
}

// AlibabaIdleParttimeSynclogAPIResponseModel is 兼职同步日志 成功返回结果
type AlibabaIdleParttimeSynclogAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_parttime_synclog_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleParttimeSynclogResult `json:"result,omitempty" xml:"result,omitempty"`
}
