package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticsShippingMethodQueryAPIResponse query shipping method API返回值
// aliexpress.local.logistics.shipping.method.query
//
// query shipping method
type AliexpressLocalLogisticsShippingMethodQueryAPIResponse struct {
	model.CommonResponse
	AliexpressLocalLogisticsShippingMethodQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressLocalLogisticsShippingMethodQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressLocalLogisticsShippingMethodQueryAPIResponseModel).Reset()
}

// AliexpressLocalLogisticsShippingMethodQueryAPIResponseModel is query shipping method 成功返回结果
type AliexpressLocalLogisticsShippingMethodQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_local_logistics_shipping_method_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// is success
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// error message
	ErrorResultMessage string `json:"error_result_message,omitempty" xml:"error_result_message,omitempty"`
	// error code
	ErrorResultCode string `json:"error_result_code,omitempty" xml:"error_result_code,omitempty"`
	// method data
	Data *QueryShippingMethodResponseDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressLocalLogisticsShippingMethodQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = ""
	m.ErrorResultMessage = ""
	m.ErrorResultCode = ""
	m.Data = nil
}

var poolAliexpressLocalLogisticsShippingMethodQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressLocalLogisticsShippingMethodQueryAPIResponse)
	},
}

// GetAliexpressLocalLogisticsShippingMethodQueryAPIResponse 从 sync.Pool 获取 AliexpressLocalLogisticsShippingMethodQueryAPIResponse
func GetAliexpressLocalLogisticsShippingMethodQueryAPIResponse() *AliexpressLocalLogisticsShippingMethodQueryAPIResponse {
	return poolAliexpressLocalLogisticsShippingMethodQueryAPIResponse.Get().(*AliexpressLocalLogisticsShippingMethodQueryAPIResponse)
}

// ReleaseAliexpressLocalLogisticsShippingMethodQueryAPIResponse 将 AliexpressLocalLogisticsShippingMethodQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressLocalLogisticsShippingMethodQueryAPIResponse(v *AliexpressLocalLogisticsShippingMethodQueryAPIResponse) {
	v.Reset()
	poolAliexpressLocalLogisticsShippingMethodQueryAPIResponse.Put(v)
}
