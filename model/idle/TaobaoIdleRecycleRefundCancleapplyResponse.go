package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收取消退款申请V2 APIResponse
taobao.idle.recycle.refund.cancleapply

回收商的回收订单取消退款申请
*/
type TaobaoIdleRecycleRefundCancleapplyAPIResponse struct {
    model.CommonResponse
    TaobaoIdleRecycleRefundCancleapplyResponse
}

type TaobaoIdleRecycleRefundCancleapplyResponse struct {
    XMLName xml.Name `xml:"idle_recycle_refund_cancleapply_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 撤销申请结果
    
    Result   *IdleTopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
