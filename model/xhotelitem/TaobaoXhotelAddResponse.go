package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店新增接口（ID重复自动更新） APIResponse
taobao.xhotel.add

添加酒店或更新酒店
*/
type TaobaoXhotelAddAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelAddResponse
}

type TaobaoXhotelAddResponse struct {
    XMLName xml.Name `xml:"xhotel_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 酒店信息
    
    Xhotel   *XHotel `json:"xhotel,omitempty" xml:"xhotel,omitempty"`

    
}
