package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退票通知 APIRequest
taobao.train.agent.returnticket.confirm.vtwo

火车票代理商接口——退票通知回调
*/
type TaobaoTrainAgentReturnticketConfirmVtwoRequest struct {
    model.Params

    // 用户id
    buyerId   int64 

    // 是否同意退票
    agreeReturn   bool 

    // 退票金额，单位分
    refundFee   int64 

    // 淘宝主订单id
    mainBizOrderId   int64 

    // 代理商id
    agentId   int64 

    // 拒绝退票原因
    refuseReturnReason   string 

    // 淘宝子订单id
    subBizOrderId   int64 

    // 关闭通知用户 true:关闭 false:不关闭
    closeRefundNotify   bool 

}

func NewTaobaoTrainAgentReturnticketConfirmVtwoRequest() *TaobaoTrainAgentReturnticketConfirmVtwoRequest{
    return &TaobaoTrainAgentReturnticketConfirmVtwoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.returnticket.confirm.vtwo"
}

func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetBuyerId(buyerId int64) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetBuyerId() int64 {
    return r.buyerId
}

func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetAgreeReturn(agreeReturn bool) error {
    r.agreeReturn = agreeReturn
    r.Set("agree_return", agreeReturn)
    return nil
}

func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetAgreeReturn() bool {
    return r.agreeReturn
}

func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetRefundFee(refundFee int64) error {
    r.refundFee = refundFee
    r.Set("refund_fee", refundFee)
    return nil
}

func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetRefundFee() int64 {
    return r.refundFee
}

func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetMainBizOrderId(mainBizOrderId int64) error {
    r.mainBizOrderId = mainBizOrderId
    r.Set("main_biz_order_id", mainBizOrderId)
    return nil
}

func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetMainBizOrderId() int64 {
    return r.mainBizOrderId
}

func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetAgentId() int64 {
    return r.agentId
}

func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetRefuseReturnReason(refuseReturnReason string) error {
    r.refuseReturnReason = refuseReturnReason
    r.Set("refuse_return_reason", refuseReturnReason)
    return nil
}

func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetRefuseReturnReason() string {
    return r.refuseReturnReason
}

func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetSubBizOrderId(subBizOrderId int64) error {
    r.subBizOrderId = subBizOrderId
    r.Set("sub_biz_order_id", subBizOrderId)
    return nil
}

func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetSubBizOrderId() int64 {
    return r.subBizOrderId
}

func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetCloseRefundNotify(closeRefundNotify bool) error {
    r.closeRefundNotify = closeRefundNotify
    r.Set("close_refund_notify", closeRefundNotify)
    return nil
}

func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetCloseRefundNotify() bool {
    return r.closeRefundNotify
}

