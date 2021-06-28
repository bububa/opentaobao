package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建主动预约开通条件 APIResponse
tmall.servicecenter.reservecond.create

1、设置主动预约开通条件
*/
type TmallServicecenterReservecondCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_reservecond_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    MsgSuccess   bool `json:"msg_success,omitempty" xml:"