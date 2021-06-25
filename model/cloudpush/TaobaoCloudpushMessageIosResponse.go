package cloudpush

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川云推送发送消息给ios APIResponse
taobao.cloudpush.message.ios

百川云推送发送消息给iOS设备.
*/
type TaobaoCloudpushMessageIosAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCloudpushMessageIosResponse `json:"taobao_cloudpush_message_ios_response,omitempty"`
}

type TaobaoCloudpushMessageIosResponse struct {

    // 请求出现错误的错误代码.
    RequestErrorCode   int64 `json:"request_error_code,omitempty"`

    // 请求失败时候的错误信息.
    RequestErrorMsg   string `json:"request_error_msg,omitempty"`

    // 请求是否成功.
    IsSuccess   bool `json:"is_success,omitempty"`

}
