package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerUpdateppwAPIResponse 修改支付密码 API返回值
// alibaba.alsc.crm.customer.updateppw
//
// 修改支付密码
type AlibabaAlscCrmCustomerUpdateppwAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCustomerUpdateppwAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCustomerUpdateppwAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCustomerUpdateppwAPIResponseModel).Reset()
}

// AlibabaAlscCrmCustomerUpdateppwAPIResponseModel is 修改支付密码 成功返回结果
type AlibabaAlscCrmCustomerUpdateppwAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_customer_updateppw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCustomerUpdateppwAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCustomerUpdateppwAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCustomerUpdateppwAPIResponse)
	},
}

// GetAlibabaAlscCrmCustomerUpdateppwAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCustomerUpdateppwAPIResponse
func GetAlibabaAlscCrmCustomerUpdateppwAPIResponse() *AlibabaAlscCrmCustomerUpdateppwAPIResponse {
	return poolAlibabaAlscCrmCustomerUpdateppwAPIResponse.Get().(*AlibabaAlscCrmCustomerUpdateppwAPIResponse)
}

// ReleaseAlibabaAlscCrmCustomerUpdateppwAPIResponse 将 AlibabaAlscCrmCustomerUpdateppwAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCustomerUpdateppwAPIResponse(v *AlibabaAlscCrmCustomerUpdateppwAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCustomerUpdateppwAPIResponse.Put(v)
}
