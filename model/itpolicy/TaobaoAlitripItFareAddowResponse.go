package itpolicy

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票自有政策】单条单程添加 APIResponse
taobao.alitrip.it.fare.addow

自有政策单程添加接口，重复的老数据会被删除，重复判断规则同excel
*/
type TaobaoAlitripItFareAddowAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripItFareAddowResponse
}

type TaobaoAlitripItFareAddowResponse struct {
    XMLName xml.Name `xml:"alitrip_it_fare_addow_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 运价id
    
    FareId   int64 `json:"fare_id,omitempty" xml:"fare_id,omitempty"`

    
    // json格式的字符串，扩展属性，预留
    
    ExtendAttributes   string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`

    
}
