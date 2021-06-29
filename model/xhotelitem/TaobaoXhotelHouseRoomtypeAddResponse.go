package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
民宿房型信息添加 APIResponse
taobao.xhotel.house.roomtype.add

房型添加或更新
*/
type TaobaoXhotelHouseRoomtypeAddAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelHouseRoomtypeAddResponse
}

type TaobaoXhotelHouseRoomtypeAddResponse struct {
    XMLName xml.Name `xml:"xhotel_house_roomtype_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 房型信息
    
    Xroomtype   *XRoomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`

    
}
