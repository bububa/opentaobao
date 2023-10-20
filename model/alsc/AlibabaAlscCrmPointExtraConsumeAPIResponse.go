package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointExtraConsumeAPIResponse 积分补扣 API返回值
// alibaba.alsc.crm.point.extra.consume
//
// 积分补扣
type AlibabaAlscCrmPointExtraConsumeAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmPointExtraConsumeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmPointExtraConsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmPointExtraConsumeAPIResponseModel).Reset()
}

// AlibabaAlscCrmPointExtraConsumeAPIResponseModel is 积分补扣 成功返回结果
type AlibabaAlscCrmPointExtraConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_point_extra_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmPointExtraConsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmPointExtraConsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmPointExtraConsumeAPIResponse)
	},
}

// GetAlibabaAlscCrmPointExtraConsumeAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmPointExtraConsumeAPIResponse
func GetAlibabaAlscCrmPointExtraConsumeAPIResponse() *AlibabaAlscCrmPointExtraConsumeAPIResponse {
	return poolAlibabaAlscCrmPointExtraConsumeAPIResponse.Get().(*AlibabaAlscCrmPointExtraConsumeAPIResponse)
}

// ReleaseAlibabaAlscCrmPointExtraConsumeAPIResponse 将 AlibabaAlscCrmPointExtraConsumeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmPointExtraConsumeAPIResponse(v *AlibabaAlscCrmPointExtraConsumeAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmPointExtraConsumeAPIResponse.Put(v)
}
