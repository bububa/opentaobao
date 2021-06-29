package itpolicy

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票自有政策】单条删除 APIResponse
taobao.alitrip.it.fare.delete

自有政策删除接口，可以根据fareId或outId删除，根据outId删除时，如果outId不唯一，返回失败
*/
type TaobaoAlitripItFareDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripItFareDeleteResponse
}

type TaobaoAlitripItFareDeleteResponse struct {
    XMLName xml.Name `xml:"alitrip_it_fare_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // json格式的字符串，扩展属性，预留
    
    ExtendAttributes   string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`

    
}
