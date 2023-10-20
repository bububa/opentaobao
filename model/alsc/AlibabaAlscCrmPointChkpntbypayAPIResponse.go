package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointChkpntbypayAPIResponse 校验支付链路中的积分抵扣是否合法 API返回值
// alibaba.alsc.crm.point.chkpntbypay
//
// 校验支付链路中的积分抵扣是否合法
type AlibabaAlscCrmPointChkpntbypayAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmPointChkpntbypayAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmPointChkpntbypayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmPointChkpntbypayAPIResponseModel).Reset()
}

// AlibabaAlscCrmPointChkpntbypayAPIResponseModel is 校验支付链路中的积分抵扣是否合法 成功返回结果
type AlibabaAlscCrmPointChkpntbypayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_point_chkpntbypay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmPointChkpntbypayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmPointChkpntbypayAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmPointChkpntbypayAPIResponse)
	},
}

// GetAlibabaAlscCrmPointChkpntbypayAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmPointChkpntbypayAPIResponse
func GetAlibabaAlscCrmPointChkpntbypayAPIResponse() *AlibabaAlscCrmPointChkpntbypayAPIResponse {
	return poolAlibabaAlscCrmPointChkpntbypayAPIResponse.Get().(*AlibabaAlscCrmPointChkpntbypayAPIResponse)
}

// ReleaseAlibabaAlscCrmPointChkpntbypayAPIResponse 将 AlibabaAlscCrmPointChkpntbypayAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmPointChkpntbypayAPIResponse(v *AlibabaAlscCrmPointChkpntbypayAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmPointChkpntbypayAPIResponse.Put(v)
}
