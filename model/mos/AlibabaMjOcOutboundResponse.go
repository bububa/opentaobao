package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
零售商品发货 APIResponse
alibaba.mj.oc.outbound

用于接收发货的数据
*/
type AlibabaMjOcOutboundAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMjOcOutboundResponse `json:"alibaba_mj_oc_outbound_response,omitempty"` 
    AlibabaMjOcOutboundResponse
}

/* model for simplify = false
type AlibabaMjOcOutboundResponse struct {

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type AlibabaMjOcOutboundResponse struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
