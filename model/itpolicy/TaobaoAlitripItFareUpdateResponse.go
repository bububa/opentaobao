package itpolicy

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票自有政策】单条修改 APIResponse
taobao.alitrip.it.fare.update

自有政策修改接口，可以根据fareId或outId修改，outId不唯一时，不能用outId修改。当外部政策id、出发城市、到达城市、出票航司任一有变化，或往返时是否允许混舱、文件编号、可混文件编号任一有变化，将删除老数据，产生一条新政策。
*/
type TaobaoAlitripItFareUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripItFareUpdateResponse
}

type TaobaoAlitripItFareUpdateResponse struct {
    XMLName xml.Name `xml:"alitrip_it_fare_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 运价id，根据更新的内容，此id可能会重新生成
    
    FareId   int64 `json:"fare_id,omitempty" xml:"fare_id,omitempty"`

    
    // json格式的字符串，扩展属性，预留
    
    ExtendAttributes   string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`

    
}
