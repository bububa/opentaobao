package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零售商品发货 APIResponse
alibaba.mj.oc.outbound

用于接收发货的数据
*/
type AlibabaMjOcOutboundAPIResponse struct {
    model.CommonResponse
    AlibabaMjOcOutboundResponse
}

type AlibabaMjOcOutboundResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_oc_outbound_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
