package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
注册加密账号 APIResponse
taobao.top.secret.register

提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
*/
type TaobaoTopSecretRegisterAPIResponse struct {
    model.CommonResponse
    TaobaoTopSecretRegisterResponse
}

type TaobaoTopSecretRegisterResponse struct {
    XMLName xml.Name `xml:"top_secret_register_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回操作是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
