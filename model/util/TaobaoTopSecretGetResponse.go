package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取TOP通道解密秘钥 APIResponse
taobao.top.secret.get

top sdk通过api获取对应解密秘钥
*/
type TaobaoTopSecretGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"top_secret_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 下次更新秘钥间隔，单位（秒）
    
    Interval   int64 `json:"interval,omitempty" xml:"