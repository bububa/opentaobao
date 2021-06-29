package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
房型新增接口（ID重复变更新） APIResponse
taobao.xhotel.roomtype.add

房型添加或更新
*/
type TaobaoXhotelRoomtypeAddAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRoomtypeAddResponse
}

type TaobaoXhotelRoomtypeAddResponse struct {
    XMLName xml.Name `xml:"xhotel_roomtype_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 房型信息
    
    Xroomtype   *XRoomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`

    
}
