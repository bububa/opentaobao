package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫逆向纠纷查询 APIResponse
tmall.dispute.receive.get

展示商家所有退款信息
*/
type TmallDisputeReceiveGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_dispute_receive_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TmallDisputeReceiveGetResultSet `json:"result,omitempty" xml:"