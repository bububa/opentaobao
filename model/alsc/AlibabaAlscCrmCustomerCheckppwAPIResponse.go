package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerCheckppwAPIResponse 校验支付密码 API返回值
// alibaba.alsc.crm.customer.checkppw
//
// 校验支付密码
type AlibabaAlscCrmCustomerCheckppwAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCustomerCheckppwAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCustomerCheckppwAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCustomerCheckppwAPIResponseModel).Reset()
}

// AlibabaAlscCrmCustomerCheckppwAPIResponseModel is 校验支付密码 成功返回结果
type AlibabaAlscCrmCustomerCheckppwAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_customer_checkppw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCustomerCheckppwAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCustomerCheckppwAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCustomerCheckppwAPIResponse)
	},
}

// GetAlibabaAlscCrmCustomerCheckppwAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCustomerCheckppwAPIResponse
func GetAlibabaAlscCrmCustomerCheckppwAPIResponse() *AlibabaAlscCrmCustomerCheckppwAPIResponse {
	return poolAlibabaAlscCrmCustomerCheckppwAPIResponse.Get().(*AlibabaAlscCrmCustomerCheckppwAPIResponse)
}

// ReleaseAlibabaAlscCrmCustomerCheckppwAPIResponse 将 AlibabaAlscCrmCustomerCheckppwAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCustomerCheckppwAPIResponse(v *AlibabaAlscCrmCustomerCheckppwAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCustomerCheckppwAPIResponse.Put(v)
}
