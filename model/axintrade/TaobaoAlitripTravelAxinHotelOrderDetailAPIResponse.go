package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderDetailAPIResponse 阿信酒店分销-订单详情接口 API返回值
// taobao.alitrip.travel.axin.hotel.order.detail
//
// 阿信酒店订单详情的读取接口
type TaobaoAlitripTravelAxinHotelOrderDetailAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelOrderDetailAPIResponseModel
}

// TaobaoAlitripTravelAxinHotelOrderDetailAPIResponseModel is 阿信酒店分销-订单详情接口 成功返回结果
type TaobaoAlitripTravelAxinHotelOrderDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_order_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoAlitripTravelAxinHotelOrderDetailResult `json:"result,omitempty" xml:"result,omitempty"`
}
