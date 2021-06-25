package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取待退票的订单v2--增加鉴权校验 APIResponse
taobao.train.agent.returnorders.get.vtwo

代理商用来获取待退票的订单列表及数量，防止代理商掉单。
*/
type TaobaoTrainAgentReturnordersGetVtwoAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTrainAgentReturnordersGetVtwoResponse `json:"taobao_train_agent_returnorders_get_vtwo_response,omitempty"`
}

type TaobaoTrainAgentReturnordersGetVtwoResponse struct {

    // 待退票的订单数
    OrderCount   int64 `json:"order_count,omitempty"`

    // 子订单号字符串，每个订单以逗号分隔
    OrderIds   string `json:"order_ids,omitempty"`

    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 主订单id列表
    MainOrderIds   string `json:"main_order_ids,omitempty"`

    // 申请时间，每个时间以逗号分隔
    RefundApplyTimes   string `json:"refund_apply_times,omitempty"`

}
