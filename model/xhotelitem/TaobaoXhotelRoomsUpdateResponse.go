package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
房型共享库存推送接口（批量全量） APIResponse
taobao.xhotel.rooms.update

此接口用于更新多个集市酒店商品房态信息，根据传入的gids更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在淘宝集市酒店系统中不存在，则会返回错误提示。是全量更新，非增量，会把之前的房态进行覆盖。
*/
type TaobaoXhotelRoomsUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRoomsUpdateResponse
}

type TaobaoXhotelRoomsUpdateResponse struct {
    XMLName xml.Name `xml:"xhotel_rooms_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成功的gids LIST
    
    Gids   []string `json:"gids,omitempty" xml:"gids>string,omitempty"`
    
    
}
