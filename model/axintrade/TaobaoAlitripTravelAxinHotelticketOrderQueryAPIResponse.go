package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse 阿信度假交易订单查询接口 API返回值
// taobao.alitrip.travel.axin.hotelticket.order.query
//
// 阿信度假交易订单查询接口
type TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponseModel
}

// TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponseModel is 阿信度假交易订单查询接口 成功返回结果
type TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotelticket_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
