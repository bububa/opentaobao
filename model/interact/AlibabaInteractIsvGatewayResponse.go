package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
isv调用gateway APIResponse
alibaba.interact.isv.gateway

isv能够调用jae本身的server
*/
type AlibabaInteractIsvGatewayAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_isv_gateway_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // ret=0
    
    Ret   string `json:"ret,omitempty" xml:"