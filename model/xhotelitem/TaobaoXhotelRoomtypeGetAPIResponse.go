package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
房型查询接口 API返回值 
taobao.xhotel.roomtype.get

房型查询房型查询接口返回结果增加date_confirm字段
*/
type TaobaoXhotelRoomtypeGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRoomtypeGetAPIResponseModel
}

// 房型查询接口 成功返回结果
type TaobaoXhotelRoomtypeGetAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_roomtype_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询得到的RoomType
    Xroomtype   *XRoomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`
}
