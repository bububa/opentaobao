package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerUpdateAPIResponse 更新顾客信息 API返回值
// alibaba.alsc.crm.customer.update
//
// 更新顾客信息
type AlibabaAlscCrmCustomerUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCustomerUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCustomerUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCustomerUpdateAPIResponseModel).Reset()
}

// AlibabaAlscCrmCustomerUpdateAPIResponseModel is 更新顾客信息 成功返回结果
type AlibabaAlscCrmCustomerUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_customer_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCustomerUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCustomerUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCustomerUpdateAPIResponse)
	},
}

// GetAlibabaAlscCrmCustomerUpdateAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCustomerUpdateAPIResponse
func GetAlibabaAlscCrmCustomerUpdateAPIResponse() *AlibabaAlscCrmCustomerUpdateAPIResponse {
	return poolAlibabaAlscCrmCustomerUpdateAPIResponse.Get().(*AlibabaAlscCrmCustomerUpdateAPIResponse)
}

// ReleaseAlibabaAlscCrmCustomerUpdateAPIResponse 将 AlibabaAlscCrmCustomerUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCustomerUpdateAPIResponse(v *AlibabaAlscCrmCustomerUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCustomerUpdateAPIResponse.Put(v)
}
