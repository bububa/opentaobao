package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomGetAPIResponse room实体查询 API返回值
// taobao.xhotel.room.get
//
// 此接口用于查询一个商品，根据传入的gid查询商品信息。卖家只能查询自己的商品。
type TaobaoXhotelRoomGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRoomGetAPIResponseModel
}

// TaobaoXhotelRoomGetAPIResponseModel is room实体查询 成功返回结果
type TaobaoXhotelRoomGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_room_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 房间信息
	Room *FirstResult `json:"room,omitempty" xml:"room,omitempty"`
}
