package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelFscHotelListQueryAPIResponse 标准酒店信息查询-供销平台 API返回值
// taobao.alitrip.travel.fsc.hotel.list.query
//
// 供销平台标准酒店信息列表查询
type TaobaoAlitripTravelFscHotelListQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelFscHotelListQueryAPIResponseModel
}

// TaobaoAlitripTravelFscHotelListQueryAPIResponseModel is 标准酒店信息查询-供销平台 成功返回结果
type TaobaoAlitripTravelFscHotelListQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_hotel_list_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口应答对象
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
