package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelPriceQueryAPIResponse 阿信酒店分销-实时报价查询 API返回值
// taobao.alitrip.travel.axin.hotel.price.query
//
// 酒店实时报价查询
type TaobaoAlitripTravelAxinHotelPriceQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelPriceQueryAPIResponseModel
}

// TaobaoAlitripTravelAxinHotelPriceQueryAPIResponseModel is 阿信酒店分销-实时报价查询 成功返回结果
type TaobaoAlitripTravelAxinHotelPriceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_price_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回模型
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
