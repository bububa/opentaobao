package cloudpush

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川云推送发送通知给android APIResponse
taobao.cloudpush.notice.android

百川云推送发送通知给android
*/
type TaobaoCloudpushNoticeAndroidAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cloudpush_notice_android_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"