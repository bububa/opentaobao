package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
网关一次性token获取 APIResponse
taobao.top.once.token.get

网关一次性token获取
*/
type TaobaoTopOnceTokenGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"top_once_token_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // token
    
    Token   string `json:"token,omitempty" xml:"