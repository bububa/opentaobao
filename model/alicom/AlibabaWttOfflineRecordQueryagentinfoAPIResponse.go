package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWttOfflineRecordQueryagentinfoAPIResponse 线下推广充送等业务订单来源 API返回值
// alibaba.wtt.offline.record.queryagentinfo
//
// 线下推广充送等业务订单来源的查询接口
type AlibabaWttOfflineRecordQueryagentinfoAPIResponse struct {
	model.CommonResponse
	AlibabaWttOfflineRecordQueryagentinfoAPIResponseModel
}

// AlibabaWttOfflineRecordQueryagentinfoAPIResponseModel is 线下推广充送等业务订单来源 成功返回结果
type AlibabaWttOfflineRecordQueryagentinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wtt_offline_record_queryagentinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
