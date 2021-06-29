package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
复杂房价查询接口 APIResponse
taobao.xhotel.multiplerate.get

查询复杂房价，支持通过入住人数，连住天数，商品信息，房价信息查询
*/
type TaobaoXhotelMultiplerateGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelMultiplerateGetResponse
}

type TaobaoXhotelMultiplerateGetResponse struct {
    XMLName xml.Name `xml:"xhotel_multiplerate_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 复杂价格返回结果类
    
    Rates   []MultipleRate `json:"rates,omitempty" xml:"rates>multiple_rate,omitempty"`
    
    
}
