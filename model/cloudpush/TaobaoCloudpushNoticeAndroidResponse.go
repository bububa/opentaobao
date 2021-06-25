package cloudpush

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川云推送发送通知给android APIResponse
taobao.cloudpush.notice.android

百川云推送发送通知给android
*/
type TaobaoCloudpushNoticeAndroidAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCloudpushNoticeAndroidResponse `json:"taobao_cloudpush_notice_android_response,omitempty"`
}

type TaobaoCloudpushNoticeAndroidResponse struct {

    // 请求是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 请求的错误代码.
    RequestErrorCode   int64 `json:"request_error_code,omitempty"`

    // 请求的错误信息.
    RequestErrorMsg   string `json:"request_error_msg,omitempty"`

}
