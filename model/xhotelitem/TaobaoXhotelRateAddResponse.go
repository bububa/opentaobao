package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增专享房价 APIResponse
taobao.xhotel.rate.add

酒店产品库rate添加
*/
type TaobaoXhotelRateAddAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateAddResponse
}

type TaobaoXhotelRateAddResponse struct {
    XMLName xml.Name `xml:"xhotel_rate_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 酒店商品id-酒店rpID
    
    GidAndRpid   string `json:"gid_and_rpid,omitempty" xml:"gid_and_rpid,omitempty"`

    
    // results
    
    Results   []string `json:"results,omitempty" xml:"results>string,omitempty"`
    
    
    // warnMessage
    
    WarnMessage   string `json:"warn_message,omitempty" xml:"warn_message,omitempty"`

    
}
