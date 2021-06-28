package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
双11到店互动花呗红包获取token鉴权接口 APIResponse
alibaba.interact.login.alipayauth

双11到店互动花呗红包获取token鉴权接口
*/
type AlibabaInteractLoginAlipayauthAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_login_alipayauth_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   string `json:"result,omitempty" xml:"