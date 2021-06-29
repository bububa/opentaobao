package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三方账号获取 token APIResponse
alibaba.ailab.user.token.get

inside 设备的三方 app，通过 extId、schema 生成 token
*/
type AlibabaAilabUserTokenGetAPIResponse struct {
    model.CommonResponse
    AlibabaAilabUserTokenGetResponse
}

type AlibabaAilabUserTokenGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ailab_user_token_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // statusCode
    
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`

    
    // 随机 uuid，token 在5分钟后失效，token 在授权成功后失效；建议每次调用 api 获取最新 token
    
    Token   string `json:"token,omitempty" xml:"token,omitempty"`

    
}
