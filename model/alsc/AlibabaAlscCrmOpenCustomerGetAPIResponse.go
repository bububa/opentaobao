package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenCustomerGetAPIResponse 查询会员资产 API返回值
// alibaba.alsc.crm.open.customer.get
//
// 查询会员身份，资产等
type AlibabaAlscCrmOpenCustomerGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmOpenCustomerGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmOpenCustomerGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmOpenCustomerGetAPIResponseModel).Reset()
}

// AlibabaAlscCrmOpenCustomerGetAPIResponseModel is 查询会员资产 成功返回结果
type AlibabaAlscCrmOpenCustomerGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_open_customer_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmOpenCustomerGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmOpenCustomerGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmOpenCustomerGetAPIResponse)
	},
}

// GetAlibabaAlscCrmOpenCustomerGetAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmOpenCustomerGetAPIResponse
func GetAlibabaAlscCrmOpenCustomerGetAPIResponse() *AlibabaAlscCrmOpenCustomerGetAPIResponse {
	return poolAlibabaAlscCrmOpenCustomerGetAPIResponse.Get().(*AlibabaAlscCrmOpenCustomerGetAPIResponse)
}

// ReleaseAlibabaAlscCrmOpenCustomerGetAPIResponse 将 AlibabaAlscCrmOpenCustomerGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmOpenCustomerGetAPIResponse(v *AlibabaAlscCrmOpenCustomerGetAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmOpenCustomerGetAPIResponse.Put(v)
}
