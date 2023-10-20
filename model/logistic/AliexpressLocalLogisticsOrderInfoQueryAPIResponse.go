package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpresslocallogisticsorderinfoqueryAPIResponse query order details API返回值
// aliexpress.local.logistics.order.info.query
//
// query order details
type AliexpresslocallogisticsorderinfoqueryAPIResponse struct {
	model.CommonResponse
	AliexpresslocallogisticsorderinfoqueryAPIResponseModel
}

// AliexpresslocallogisticsorderinfoqueryAPIResponseModel is query order details 成功返回结果
type AliexpresslocallogisticsorderinfoqueryAPIResponseModel struct {
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
