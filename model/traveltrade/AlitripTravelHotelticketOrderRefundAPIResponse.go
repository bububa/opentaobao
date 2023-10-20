package traveltrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelhotelticketorderrefundAPIResponse 退款结结果通知 API返回值
// alitrip.travel.hotelticket.order.refund
//
// 退款结果通知
type AlitriptravelhotelticketorderrefundAPIResponse struct {
	model.CommonResponse
	AlitriptravelhotelticketorderrefundAPIResponseModel
}

// AlitriptravelhotelticketorderrefundAPIResponseModel is 退款结结果通知 成功返回结果
type AlitriptravelhotelticketorderrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_hotelticket_order_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
