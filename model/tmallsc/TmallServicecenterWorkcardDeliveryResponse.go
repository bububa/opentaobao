package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
开始配送工单 APIResponse
tmall.servicecenter.workcard.delivery

服务商调用该接口通知天猫服务平台服务商工人已开始配送工单
*/
type TmallServicecenterWorkcardDeliveryAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardDeliveryResponse
}

type TmallServicecenterWorkcardDeliveryResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_delivery_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    MsgSuccess   bool `json:"msg_success,omitempty" xml:"msg_success,omitempty"`

    
    // 返回信息
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 返回code
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}
