package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取订单退票信息 APIResponse
taobao.train.agent.get.refund

代理商获取订单信息回调API
*/
type TaobaoTrainAgentGetRefundAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTrainAgentGetRefundResponse `json:"train_agent_get_refund_response,omitempty"` 
    TaobaoTrainAgentGetRefundResponse
}

/* model for simplify = false
type TaobaoTrainAgentGetRefundResponse struct {

    // 系统自动生成
    
    TopRefundApplyList   string `json:"top_refund_apply_list,omitempty"`
    

}
*/

type TaobaoTrainAgentGetRefundResponse struct {

    // 系统自动生成
    TopRefundApplyList   string `json:"top_refund_apply_list,omitempty"`

}
