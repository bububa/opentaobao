package cloudpush

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川云推送发送消息给android APIResponse
taobao.cloudpush.message.android

百川用户使用云推送发送消息给android
*/
type TaobaoCloudpushMessageAndroidAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoCloudpushMessageAndroidResponse `json:"cloudpush_message_android_response,omitempty"` 
    TaobaoCloudpushMessageAndroidResponse
}

/* model for simplify = false
type TaobaoCloudpushMessageAndroidResponse struct {

    // 请求是否成功.
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 若请求失败，则返回对应的error code.
    
    RequestErrorCode   int64 `json:"request_error_code,omitempty"`
    

    // 请求失败后返回的错误信息.
    
    RequestErrorMsg   string `json:"request_error_msg,omitempty"`
    

}
*/

type TaobaoCloudpushMessageAndroidResponse struct {

    // 请求是否成功.
    IsSuccess   bool `json:"is_success,omitempty"`

    // 若请求失败，则返回对应的error code.
    RequestErrorCode   int64 `json:"request_error_code,omitempty"`

    // 请求失败后返回的错误信息.
    RequestErrorMsg   string `json:"request_error_msg,omitempty"`

}
