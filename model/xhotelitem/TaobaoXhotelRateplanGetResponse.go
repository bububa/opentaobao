package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
价格计划rateplan查询 APIResponse
taobao.xhotel.rateplan.get

酒店产品库rateplan查询
*/
type TaobaoXhotelRateplanGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateplanGetResponse
}

type TaobaoXhotelRateplanGetResponse struct {
    XMLName xml.Name `xml:"xhotel_rateplan_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // rateplan
    
    Rateplan   *RatePlan `json:"rateplan,omitempty" xml:"rateplan,omitempty"`

    
}
