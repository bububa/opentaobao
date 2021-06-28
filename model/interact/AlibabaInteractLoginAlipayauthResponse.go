package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
双11到店互动花呗红包获取token鉴权接口 APIResponse
alibaba.interact.login.alipayauth

双11到店互动花呗红包获取token鉴权接口
*/
type AlibabaInteractLoginAlipayauthAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractLoginAlipayauthResponse `json:"alibaba_interact_login_alipayauth_response,omitempty"` 
    AlibabaInteractLoginAlipayauthResponse
}

/* model for simplify = false
type AlibabaInteractLoginAlipayauthResponse struct {

    // 返回值
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractLoginAlipayauthResponse struct {

    // 返回值
    Result   string `json:"result,omitempty"`

}
