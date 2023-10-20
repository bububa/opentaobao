package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelroomlistqueryAPIResponse 阿信酒店分销-标准酒店房型列表查询 API返回值
// taobao.alitrip.travel.axin.hotel.room.list.query
//
// 标准酒店房型列表查询
type TaobaoalitriptravelaxinhotelroomlistqueryAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinhotelroomlistqueryAPIResponseModel
}

// TaobaoalitriptravelaxinhotelroomlistqueryAPIResponseModel is 阿信酒店分销-标准酒店房型列表查询 成功返回结果
type TaobaoalitriptravelaxinhotelroomlistqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_room_list_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回模型
	Result *BaseResultApiDto `json:"result,omitempty" xml:"result,omitempty"`
}
