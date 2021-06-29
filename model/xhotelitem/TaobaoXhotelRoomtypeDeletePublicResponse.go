package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家删除房型数据接口 APIResponse
taobao.xhotel.roomtype.delete.public

房型删除TOP接口
*/
type TaobaoXhotelRoomtypeDeletePublicAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRoomtypeDeletePublicResponse
}

type TaobaoXhotelRoomtypeDeletePublicResponse struct {
    XMLName xml.Name `xml:"xhotel_roomtype_delete_public_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoXhotelRoomtypeDeletePublicResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
