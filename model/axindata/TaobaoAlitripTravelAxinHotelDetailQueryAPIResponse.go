package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelDetailQueryAPIResponse 阿信酒店分销-标准酒店详情查询 API返回值
// taobao.alitrip.travel.axin.hotel.detail.query
//
// 标准酒店详情查询
type TaobaoAlitripTravelAxinHotelDetailQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelDetailQueryAPIResponseModel
}

// TaobaoAlitripTravelAxinHotelDetailQueryAPIResponseModel is 阿信酒店分销-标准酒店详情查询 成功返回结果
type TaobaoAlitripTravelAxinHotelDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回模型
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
