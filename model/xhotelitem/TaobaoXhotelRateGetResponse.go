package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店产品库rate查询 APIResponse
taobao.xhotel.rate.get

酒店产品库rate查询
*/
type TaobaoXhotelRateGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateGetResponse
}

type TaobaoXhotelRateGetResponse struct {
    XMLName xml.Name `xml:"xhotel_rate_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // rate
    
    Rate   *Rate `json:"rate,omitempty" xml:"rate,omitempty"`

    
}
