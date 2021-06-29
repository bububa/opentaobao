package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店房型与房价查询 APIResponse
taobao.xhotel.baseinfo.room.get

根据outHid/hid获取酒店的房型和价格信息
*/
type TaobaoXhotelBaseinfoRoomGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelBaseinfoRoomGetResponse
}

type TaobaoXhotelBaseinfoRoomGetResponse struct {
    XMLName xml.Name `xml:"xhotel_baseinfo_room_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoXhotelBaseinfoRoomGetResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
