package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
主动预约条件更新 APIResponse
tmall.servicecenter.reservecond.update

1、设置主动预约开通条件
*/
type TmallServicecenterReservecondUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_reservecond_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    MsgSuccess   bool `json:"msg_success,omitempty" xml:"