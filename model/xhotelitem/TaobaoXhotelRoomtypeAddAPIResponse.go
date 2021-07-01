package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
房型新增接口（ID重复变更新） API返回值 
taobao.xhotel.roomtype.add

房型添加或更新
*/
type TaobaoXhotelRoomtypeAddAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRoomtypeAddAPIResponseModel
}

// 房型新增接口（ID重复变更新） 成功返回结果
type TaobaoXhotelRoomtypeAddAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_roomtype_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 房型信息
    Xroomtype   *XRoomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`
}
