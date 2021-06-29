package idleparttime

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
兼职通知接口 APIResponse
alibaba.idle.parttime.notify

服务商侧有岗位状态变更对我们进行通知(岗位关闭, 录取状态)
*/
type AlibabaIdleParttimeNotifyAPIResponse struct {
    model.CommonResponse
    AlibabaIdleParttimeNotifyResponse
}

type AlibabaIdleParttimeNotifyResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_parttime_notify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
