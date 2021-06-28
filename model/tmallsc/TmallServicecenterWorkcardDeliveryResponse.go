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
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_workcard_delivery_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    MsgSuccess   bool `json:"msg_success,omitempty" xml:"