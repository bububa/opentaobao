package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelListGetAPIResponse 标准酒店信息查询-阿信 API返回值
// taobao.alitrip.travel.axin.hotel.list.get
//
// 阿信酒店分销基础数据查询
type TaobaoAlitripTravelAxinHotelListGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelListGetAPIResponseModel
}

// TaobaoAlitripTravelAxinHotelListGetAPIResponseModel is 标准酒店信息查询-阿信 成功返回结果
type TaobaoAlitripTravelAxinHotelListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripTravelAxinHotelListGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
