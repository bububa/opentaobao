package cloudpush

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推送通知给ios设备 APIResponse
taobao.cloudpush.notice.ios

推送通知给ios设备
*/
type TaobaoCloudpushNoticeIosAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cloudpush_notice_ios_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求是否成功.
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"