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
    TmallDisputeReceiveGetResponse
}

type TmallDisputeReceiveGetResponse struct {
    XMLName xml.Name `xml:"tmall_dispute_receive_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TmallDisputeReceiveGetResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
