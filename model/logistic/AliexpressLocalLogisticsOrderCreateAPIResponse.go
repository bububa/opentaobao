package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticsOrderCreateAPIResponse create logistics order API返回值
// aliexpress.local.logistics.order.create
//
// create logistics order
type AliexpressLocalLogisticsOrderCreateAPIResponse struct {
	model.CommonResponse
	AliexpressLocalLogisticsOrderCreateAPIResponseModel
}

// AliexpressLocalLogisticsOrderCreateAPIResponseModel is create logistics order 成功返回结果
type AliexpressLocalLogisticsOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_local_logistics_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error message
	ErrorResultMessage string `json:"error_result_message,omitempty" xml:"error_result_message,omitempty"`
	// error code
	ErrorResultCode string `json:"error_result_code,omitempty" xml:"error_result_code,omitempty"`
	// response info
	Data *AelogisticsOrderDto `json:"data,omitempty" xml:"data,omitempty"`
	// interface status
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
