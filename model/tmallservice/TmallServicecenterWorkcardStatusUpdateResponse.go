package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商反馈服务的执行情况 APIResponse
tmall.servicecenter.workcard.status.update

1 如果服务商受理了此服务，修改合同状态为：已受理=3<br/>2 如果服务商没有受理此服务，修改合同状态为：已拒绝=10<br/>3 如果服务商执行了此服务，修改合同状态为：已执行=4<br/>4 如果服务商执行服务成功，修改合同状态为：已完成=5<br/>5 如果此工单为合同类型的工单，当服务商受理了此服务后，会进行分账
*/
type TmallServicecenterWorkcardStatusUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_workcard_status_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Rs   bool `json:"rs,omitempty" xml:"