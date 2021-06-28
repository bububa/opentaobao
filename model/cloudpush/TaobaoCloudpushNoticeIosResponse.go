package cloudpush

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推送通知给ios设备 APIResponse
taobao.cloudpush.notice.ios

推送通知给ios设备
*/
type TaobaoCloudpushNoticeIosAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoCloudpushNoticeIosResponse `json:"cloudpush_notice_ios_response,omitempty"` 
    TaobaoCloudpushNoticeIosResponse
}

/* model for simplify = false
type TaobaoCloudpushNoticeIosResponse struct {

    // 请求是否成功.
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 请求错误时产生的错误代码.
    
    RequestErrorCode   int64 `json:"request_error_code,omitempty"`
    

    // 请求产生的错误信息.
    
    RequestErrorMsg   string `json:"request_error_msg,omitempty"`
    

}
*/

type TaobaoCloudpushNoticeIosResponse struct {

    // 请求是否成功.
    IsSuccess   bool `json:"is_success,omitempty"`

    // 请求错误时产生的错误代码.
    RequestErrorCode   int64 `json:"request_error_code,omitempty"`

    // 请求产生的错误信息.
    RequestErrorMsg   string `json:"request_error_msg,omitempty"`

}
