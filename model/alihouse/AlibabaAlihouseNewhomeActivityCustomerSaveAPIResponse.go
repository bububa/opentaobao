package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse 销售活动批量保存定向用户 API返回值
// alibaba.alihouse.newhome.activity.customer.save
//
// 销售活动批量保存定向用户
type AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponseModel is 销售活动批量保存定向用户 成功返回结果
type AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_activity_customer_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回报文属性
	Result *AlibabaAlihouseNewhomeActivityCustomerSaveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse
func GetAlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse() *AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse {
	return poolAlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse.Get().(*AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse 将 AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse(v *AlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeActivityCustomerSaveAPIResponse.Put(v)
}
