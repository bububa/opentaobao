package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
room实体查询 APIResponse
taobao.xhotel.room.get

此接口用于查询一个商品，根据传入的gid查询商品信息。卖家只能查询自己的商品。
*/
type TaobaoXhotelRoomGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRoomGetResponse
}

type TaobaoXhotelRoomGetResponse struct {
    XMLName xml.Name `xml:"xhotel_room_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 房间信息
    
    Room   *FirstResult `json:"room,omitempty" xml:"room,omitempty"`

    
}
