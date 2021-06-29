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
    TaobaoCloudpushMessageIosResponse
}

type TaobaoCloudpushMessageIosResponse struct {
    XMLName xml.Name `xml:"cloudpush_message_ios_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求出现错误的错误代码.
    
    RequestErrorCode   int64 `json:"request_error_code,omitempty" xml:"request_error_code,omitempty"`

    
    // 请求失败时候的错误信息.
    
    RequestErrorMsg   string `json:"request_error_msg,omitempty" xml:"request_error_msg,omitempty"`

    
    // 请求是否成功.
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
