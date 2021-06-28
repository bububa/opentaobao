package antifraud

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
人机识别 APIResponse
taobao.collinafacade.nocaptcha.sig.authenticate

人机识别颁发签名串后,本接口负责向ISV提供签名串校验服务
*/
type TaobaoCollinafacadeNocaptchaSigAuthenticateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"collinafacade_nocaptcha_sig_authenticate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 服务出参
    
    Ret   int64 `json:"ret,omitempty" xml:"