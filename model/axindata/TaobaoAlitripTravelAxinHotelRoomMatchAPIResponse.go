package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelroommatchAPIResponse 阿信酒店房型匹配 API返回值
// taobao.alitrip.travel.axin.hotel.room.match
//
// 阿信酒店房型匹配
type TaobaoalitriptravelaxinhotelroommatchAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinhotelroommatchAPIResponseModel
}

// TaobaoalitriptravelaxinhotelroommatchAPIResponseModel is 阿信酒店房型匹配 成功返回结果
type TaobaoalitriptravelaxinhotelroommatchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_room_match_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
