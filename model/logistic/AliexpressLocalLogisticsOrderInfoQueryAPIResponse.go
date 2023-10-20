package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticsOrderInfoQueryAPIResponse query order details API返回值
// aliexpress.local.logistics.order.info.query
//
// query order details
type AliexpressLocalLogisticsOrderInfoQueryAPIResponse struct {
	model.CommonResponse
	AliexpressLocalLogisticsOrderInfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressLocalLogisticsOrderInfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressLocalLogisticsOrderInfoQueryAPIResponseModel).Reset()
}

// AliexpressLocalLogisticsOrderInfoQueryAPIResponseModel is query order details 成功返回结果
type AliexpressLocalLogisticsOrderInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_local_logistics_order_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error_result_message
	ErrorResultMessage string `json:"error_result_message,omitempty" xml:"error_result_message,omitempty"`
	// error_result_code
	ErrorResultCode string `json:"error_result_code,omitempty" xml:"error_result_code,omitempty"`
	// data
	Data *OrderDto `json:"data,omitempty" xml:"data,omitempty"`
	// is_success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressLocalLogisticsOrderInfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorResultMessage = ""
	m.ErrorResultCode = ""
	m.Data = nil
	m.IsSuccess = false
}

var poolAliexpressLocalLogisticsOrderInfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressLocalLogisticsOrderInfoQueryAPIResponse)
	},
}

// GetAliexpressLocalLogisticsOrderInfoQueryAPIResponse 从 sync.Pool 获取 AliexpressLocalLogisticsOrderInfoQueryAPIResponse
func GetAliexpressLocalLogisticsOrderInfoQueryAPIResponse() *AliexpressLocalLogisticsOrderInfoQueryAPIResponse {
	return poolAliexpressLocalLogisticsOrderInfoQueryAPIResponse.Get().(*AliexpressLocalLogisticsOrderInfoQueryAPIResponse)
}

// ReleaseAliexpressLocalLogisticsOrderInfoQueryAPIResponse 将 AliexpressLocalLogisticsOrderInfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressLocalLogisticsOrderInfoQueryAPIResponse(v *AliexpressLocalLogisticsOrderInfoQueryAPIResponse) {
	v.Reset()
	poolAliexpressLocalLogisticsOrderInfoQueryAPIResponse.Put(v)
}
