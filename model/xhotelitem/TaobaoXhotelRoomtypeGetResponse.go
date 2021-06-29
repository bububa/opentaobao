package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
房型查询接口 APIResponse
taobao.xhotel.roomtype.get

房型查询房型查询接口返回结果增加date_confirm字段
*/
type TaobaoXhotelRoomtypeGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRoomtypeGetResponse
}

type TaobaoXhotelRoomtypeGetResponse struct {
    XMLName xml.Name `xml:"xhotel_roomtype_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询得到的RoomType
    
    Xroomtype   *XRoomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`

    
}
