package logistic

import (
	"encoding/xml"

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
