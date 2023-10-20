package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointExtrachargeAPIResponse 积分补录 API返回值
// alibaba.alsc.crm.point.extracharge
//
// 积分补录
type AlibabaAlscCrmPointExtrachargeAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmPointExtrachargeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmPointExtrachargeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmPointExtrachargeAPIResponseModel).Reset()
}

// AlibabaAlscCrmPointExtrachargeAPIResponseModel is 积分补录 成功返回结果
type AlibabaAlscCrmPointExtrachargeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_point_extracharge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmPointExtrachargeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmPointExtrachargeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmPointExtrachargeAPIResponse)
	},
}

// GetAlibabaAlscCrmPointExtrachargeAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmPointExtrachargeAPIResponse
func GetAlibabaAlscCrmPointExtrachargeAPIResponse() *AlibabaAlscCrmPointExtrachargeAPIResponse {
	return poolAlibabaAlscCrmPointExtrachargeAPIResponse.Get().(*AlibabaAlscCrmPointExtrachargeAPIResponse)
}

// ReleaseAlibabaAlscCrmPointExtrachargeAPIResponse 将 AlibabaAlscCrmPointExtrachargeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmPointExtrachargeAPIResponse(v *AlibabaAlscCrmPointExtrachargeAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmPointExtrachargeAPIResponse.Put(v)
}
