package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerResetppwAPIResponse 重置支付密码 API返回值
// alibaba.alsc.crm.customer.resetppw
//
// 重置支付密码
type AlibabaAlscCrmCustomerResetppwAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCustomerResetppwAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCustomerResetppwAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCustomerResetppwAPIResponseModel).Reset()
}

// AlibabaAlscCrmCustomerResetppwAPIResponseModel is 重置支付密码 成功返回结果
type AlibabaAlscCrmCustomerResetppwAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_customer_resetppw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCustomerResetppwAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCustomerResetppwAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCustomerResetppwAPIResponse)
	},
}

// GetAlibabaAlscCrmCustomerResetppwAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCustomerResetppwAPIResponse
func GetAlibabaAlscCrmCustomerResetppwAPIResponse() *AlibabaAlscCrmCustomerResetppwAPIResponse {
	return poolAlibabaAlscCrmCustomerResetppwAPIResponse.Get().(*AlibabaAlscCrmCustomerResetppwAPIResponse)
}

// ReleaseAlibabaAlscCrmCustomerResetppwAPIResponse 将 AlibabaAlscCrmCustomerResetppwAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCustomerResetppwAPIResponse(v *AlibabaAlscCrmCustomerResetppwAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCustomerResetppwAPIResponse.Put(v)
}
