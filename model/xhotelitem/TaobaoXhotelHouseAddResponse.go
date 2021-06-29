package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
非标准民宿房源添加 APIResponse
taobao.xhotel.house.add

添加酒店或更新酒店
*/
type TaobaoXhotelHouseAddAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelHouseAddResponse
}

type TaobaoXhotelHouseAddResponse struct {
    XMLName xml.Name `xml:"xhotel_house_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 酒店信息
    
    Xhotel   *XHotel `json:"xhotel,omitempty" xml:"xhotel,omitempty"`

    
}
