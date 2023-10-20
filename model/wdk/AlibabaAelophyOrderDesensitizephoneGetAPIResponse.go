package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyOrderDesensitizephoneGetAPIResponse 获取订单隐私号 API返回值
// alibaba.aelophy.order.desensitizephone.get
//
// 获取订单隐私号
type AlibabaAelophyOrderDesensitizephoneGetAPIResponse struct {
	model.CommonResponse
	AlibabaAelophyOrderDesensitizephoneGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAelophyOrderDesensitizephoneGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAelophyOrderDesensitizephoneGetAPIResponseModel).Reset()
}

// AlibabaAelophyOrderDesensitizephoneGetAPIResponseModel is 获取订单隐私号 成功返回结果
type AlibabaAelophyOrderDesensitizephoneGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aelophy_order_desensitizephone_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ApiErrCode string `json:"api_err_code,omitempty" xml:"api_err_code,omitempty"`
	// 错误信息
	ApiErrMsg string `json:"api_err_msg,omitempty" xml:"api_err_msg,omitempty"`
	// 订单隐私号信息
	Model *OrderDesensitizePhoneResult `json:"model,omitempty" xml:"model,omitempty"`
	// 调用是否成功
	ApiSuccess bool `json:"api_success,omitempty" xml:"api_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAelophyOrderDesensitizephoneGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiErrCode = ""
	m.ApiErrMsg = ""
	m.Model = nil
	m.ApiSuccess = false
}

var poolAlibabaAelophyOrderDesensitizephoneGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAelophyOrderDesensitizephoneGetAPIResponse)
	},
}

// GetAlibabaAelophyOrderDesensitizephoneGetAPIResponse 从 sync.Pool 获取 AlibabaAelophyOrderDesensitizephoneGetAPIResponse
func GetAlibabaAelophyOrderDesensitizephoneGetAPIResponse() *AlibabaAelophyOrderDesensitizephoneGetAPIResponse {
	return poolAlibabaAelophyOrderDesensitizephoneGetAPIResponse.Get().(*AlibabaAelophyOrderDesensitizephoneGetAPIResponse)
}

// ReleaseAlibabaAelophyOrderDesensitizephoneGetAPIResponse 将 AlibabaAelophyOrderDesensitizephoneGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAelophyOrderDesensitizephoneGetAPIResponse(v *AlibabaAelophyOrderDesensitizephoneGetAPIResponse) {
	v.Reset()
	poolAlibabaAelophyOrderDesensitizephoneGetAPIResponse.Put(v)
}
