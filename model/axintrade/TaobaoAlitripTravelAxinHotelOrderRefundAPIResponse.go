package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse 阿信酒店分销订单退款API API返回值
// taobao.alitrip.travel.axin.hotel.order.refund
//
// 阿信酒店分销订单退款
type TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelOrderRefundAPIResponseModel
}

// TaobaoAlitripTravelAxinHotelOrderRefundAPIResponseModel is 阿信酒店分销订单退款API 成功返回结果
type TaobaoAlitripTravelAxinHotelOrderRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_order_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
