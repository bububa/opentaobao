package mtop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取文件上传授权 APIResponse
taobao.mtop.upload.token.get

获取mtop文件上传授权
*/
type TaobaoMtopUploadTokenGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMtopUploadTokenGetResponse `json:"mtop_upload_token_get_response,omitempty"` 
    TaobaoMtopUploadTokenGetResponse
}

/* model for simplify = false
type TaobaoMtopUploadTokenGetResponse struct {

    // msg
    
    Message   string `json:"message,omitempty"`
    

    // code
    
    Code   string `json:"code,omitempty"`
    

    // 单次上传文件块最大大小，单位 byte
    
    MaxBodyLength   int64 `json:"max_body_length,omitempty"`
    

    // 单个文件重试上传次数
    
    MaxRetryTimes   int64 `json:"max_retry_times,omitempty"`
    

    // 本次指定的上传文件服务器地址
    
    ServerAddress   string `json:"server_address,omitempty"`
    

    // token失效时间点
    
    Timeout   int64 `json:"timeout,omitempty"`
    

    // 颁发的上传令牌
    
    Token   string `json:"token,omitempty"`
    

}
*/

type TaobaoMtopUploadTokenGetResponse struct {

    // msg
    Message   string `json:"message,omitempty"`

    // code
    Code   string `json:"code,omitempty"`

    // 单次上传文件块最大大小，单位 byte
    MaxBodyLength   int64 `json:"max_body_length,omitempty"`

    // 单个文件重试上传次数
    MaxRetryTimes   int64 `json:"max_retry_times,omitempty"`

    // 本次指定的上传文件服务器地址
    ServerAddress   string `json:"server_address,omitempty"`

    // token失效时间点
    Timeout   int64 `json:"timeout,omitempty"`

    // 颁发的上传令牌
    Token   string `json:"token,omitempty"`

}
