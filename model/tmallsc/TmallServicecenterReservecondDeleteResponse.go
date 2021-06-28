package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除主动预约开通条件 APIResponse
tmall.servicecenter.reservecond.delete

删除主动预约开通条件
*/
type TmallServicecenterReservecondDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_reservecond_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    MsgSuccess   bool `json:"msg_success,omitempty" xml:"