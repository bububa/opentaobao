package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelorderpayAPIResponse 阿信酒店分销订单支付 API返回值
// taobao.alitrip.travel.axin.hotel.order.pay
//
// 阿信酒店分销订单支付
type TaobaoalitriptravelaxinhotelorderpayAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinhotelorderpayAPIResponseModel
}

// TaobaoalitriptravelaxinhotelorderpayAPIResponseModel is 阿信酒店分销订单支付 成功返回结果
type TaobaoalitriptravelaxinhotelorderpayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_order_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoalitriptravelaxinhotelorderpayResult `json:"result,omitempty" xml:"result,omitempty"`
}
