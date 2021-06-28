package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
网关一次性token获取 APIResponse
taobao.top.once.token.get

网关一次性token获取
*/
type TaobaoTopOnceTokenGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTopOnceTokenGetResponse `json:"top_once_token_get_response,omitempty"` 
    TaobaoTopOnceTokenGetResponse
}

/* model for simplify = false
type TaobaoTopOnceTokenGetResponse struct {

    // token
    
    Token   string `json:"token,omitempty"`
    

    // 响应编码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 失败详情
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

}
*/

type TaobaoTopOnceTokenGetResponse struct {

    // token
    Token   string `json:"token,omitempty"`

    // 响应编码
    ResultCode   string `json:"result_code,omitempty"`

    // 失败详情
    ResultMsg   string `json:"result_msg,omitempty"`

}
