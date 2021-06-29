package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
房型库存推送接口（全量推送） APIResponse
taobao.xhotel.room.update

此接口用于更新一个酒店商品，根据传入的gid更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在酒店系统中不存在，则会返回错误提示。
*/
type TaobaoXhotelRoomUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRoomUpdateResponse
}

type TaobaoXhotelRoomUpdateResponse struct {
    XMLName xml.Name `xml:"xhotel_room_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // gid酒店商品id
    
    Gid   int64 `json:"gid,omitempty" xml:"gid,omitempty"`

    
}
