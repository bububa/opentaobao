package cloudpush

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川用户使用云推送高级推送接口 APIResponse
taobao.cloudpush.push

百川用户使用云推送高级推送接口
*/
type TaobaoCloudpushPushAPIResponse struct {
    model.CommonResponse
    TaobaoCloudpushPushResponse
}

type TaobaoCloudpushPushResponse struct {
    XMLName xml.Name `xml:"cloudpush_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 请求失败对应的错误代码.
    
    RequestErrorCode   int64 `json:"request_error_code,omitempty" xml:"request_error_code,omitempty"`

    
    // 请求失败的错误信息.
    
    RequestErrorMsg   string `json:"request_error_msg,omitempty" xml:"request_error_msg,omitempty"`

    
}
