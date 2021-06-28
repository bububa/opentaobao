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
	RequestId     string         `json:"request_id,omitempty" xml:"top_secret_register_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回操作是否成功
    
    Result   bool `json:"result,omitempty" xml:"