package cloudpush

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川云推送发送消息给ios APIResponse
taobao.cloudpush.message.ios

百川云推送发送消息给iOS设备.
*/
type TaobaoCloudpushMessageIosAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cloudpush_message_ios_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求出现错误的错误代码.
    
    RequestErrorCode   int64 `json:"request_error_code,omitempty" xml:"