package traveltrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelhotelticketordercreateAPIResponse 创单(支付订单)通知 API返回值
// alitrip.travel.hotelticket.order.create
//
// 创单(支付订单)通知
type AlitriptravelhotelticketordercreateAPIResponse struct {
	model.CommonResponse
	AlitriptravelhotelticketordercreateAPIResponseModel
}

// AlitriptravelhotelticketordercreateAPIResponseModel is 创单(支付订单)通知 成功返回结果
type AlitriptravelhotelticketordercreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_hotelticket_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
