package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomsIncrementAPIResponse 房型库存推送接口（批量增量） API返回值
// taobao.xhotel.rooms.increment
//
// Room库存增量更新接口，用户仅需要更新ROOM库存中发生变化的库存日历即可。
type TaobaoXhotelRoomsIncrementAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRoomsIncrementAPIResponseModel
}

// TaobaoXhotelRoomsIncrementAPIResponseModel is 房型库存推送接口（批量增量） 成功返回结果
type TaobaoXhotelRoomsIncrementAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rooms_increment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功的gids LIST
	Gids []string `json:"gids,omitempty" xml:"gids>string,omitempty"`
}
