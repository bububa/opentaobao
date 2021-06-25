package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
isv调用gateway APIResponse
alibaba.interact.isv.gateway

isv能够调用jae本身的server
*/
type AlibabaInteractIsvGatewayAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractIsvGatewayResponse `json:"alibaba_interact_isv_gateway_response,omitempty"`
}

type AlibabaInteractIsvGatewayResponse struct {

    // ret=0
    Ret   string `json:"ret,omitempty"`

}
