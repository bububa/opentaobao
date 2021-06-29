package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店产品库rateplan添加 APIResponse
taobao.xhotel.rateplan.add

酒店产品库rateplan
*/
type TaobaoXhotelRateplanAddAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateplanAddResponse
}

type TaobaoXhotelRateplanAddResponse struct {
    XMLName xml.Name `xml:"xhotel_rateplan_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 生成的rp id
    
    Rpid   int64 `json:"rpid,omitempty" xml:"rpid,omitempty"`

    
}
