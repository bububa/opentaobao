package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelShidListQueryAPIResponse 阿信酒店分销-标准酒店id列表查询 API返回值
// taobao.alitrip.travel.axin.hotel.shid.list.query
//
// 标准酒店id列表查询
type TaobaoAlitripTravelAxinHotelShidListQueryAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelShidListQueryAPIResponseModel
}

// TaobaoAlitripTravelAxinHotelShidListQueryAPIResponseModel is 阿信酒店分销-标准酒店id列表查询 成功返回结果
type TaobaoAlitripTravelAxinHotelShidListQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_shid_list_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 应答模型
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
