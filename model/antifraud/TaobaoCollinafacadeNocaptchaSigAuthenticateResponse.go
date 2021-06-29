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
    TaobaoCollinafacadeNocaptchaSigAuthenticateResponse
}

type TaobaoCollinafacadeNocaptchaSigAuthenticateResponse struct {
    XMLName xml.Name `xml:"collinafacade_nocaptcha_sig_authenticate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 服务出参
    
    Ret   int64 `json:"ret,omitempty" xml:"ret,omitempty"`

    
    // 返回authenticateResult
    
    RetDetail   *SigAuthenticateResult `json:"ret_detail,omitempty" xml:"ret_detail,omitempty"`

    
}
