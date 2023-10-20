package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderPayAPIResponse 阿信酒店分销订单支付 API返回值
// taobao.alitrip.travel.axin.hotel.order.pay
//
// 阿信酒店分销订单支付
type TaobaoAlitripTravelAxinHotelOrderPayAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelOrderPayAPIResponseModel
}

// TaobaoAlitripTravelAxinHotelOrderPayAPIResponseModel is 阿信酒店分销订单支付 成功返回结果
type TaobaoAlitripTravelAxinHotelOrderPayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_order_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripTravelAxinHotelOrderPayResult `json:"result,omitempty" xml:"result,omitempty"`
}
