package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvGoosefishOrderCreateAPIResponse 闲鱼三方安康容器订单创建 API返回值
// alibaba.idle.isv.goosefish.order.create
//
// 闲鱼三方安康容器订单创建
type AlibabaIdleIsvGoosefishOrderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvGoosefishOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvGoosefishOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvGoosefishOrderCreateAPIResponseModel).Reset()
}

// AlibabaIdleIsvGoosefishOrderCreateAPIResponseModel is 闲鱼三方安康容器订单创建 成功返回结果
type AlibabaIdleIsvGoosefishOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_goosefish_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ApiErrorCode string `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
	// 错误信息
	ApiErrorMsg string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
	// 创单结果
	Data *OrderCreateResult `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	ApiSuccess bool `json:"api_success,omitempty" xml:"api_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvGoosefishOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiErrorCode = ""
	m.ApiErrorMsg = ""
	m.Data = nil
	m.ApiSuccess = false
}

var poolAlibabaIdleIsvGoosefishOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvGoosefishOrderCreateAPIResponse)
	},
}

// GetAlibabaIdleIsvGoosefishOrderCreateAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvGoosefishOrderCreateAPIResponse
func GetAlibabaIdleIsvGoosefishOrderCreateAPIResponse() *AlibabaIdleIsvGoosefishOrderCreateAPIResponse {
	return poolAlibabaIdleIsvGoosefishOrderCreateAPIResponse.Get().(*AlibabaIdleIsvGoosefishOrderCreateAPIResponse)
}

// ReleaseAlibabaIdleIsvGoosefishOrderCreateAPIResponse 将 AlibabaIdleIsvGoosefishOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvGoosefishOrderCreateAPIResponse(v *AlibabaIdleIsvGoosefishOrderCreateAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvGoosefishOrderCreateAPIResponse.Put(v)
}
