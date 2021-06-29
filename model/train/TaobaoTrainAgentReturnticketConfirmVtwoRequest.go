package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退票通知 API请求
taobao.train.agent.returnticket.confirm.vtwo

火车票代理商接口——退票通知回调
*/
type TaobaoTrainAgentReturnticketConfirmVtwoRequest struct {
    model.Params
    // 用户id
    _buyerId   int64
    // 是否同意退票
    _agreeReturn   bool
    // 退票金额，单位分
    _refundFee   int64
    // 淘宝主订单id
    _mainBizOrderId   int64
    // 代理商id
    _agentId   int64
    // 拒绝退票原因
    _refuseReturnReason   string
    // 淘宝子订单id
    _subBizOrderId   int64
    // 关闭通知用户 true:关闭 false:不关闭
    _closeRefundNotify   bool
}

// 初始化TaobaoTrainAgentReturnticketConfirmVtwoRequest对象
func NewTaobaoTrainAgentReturnticketConfirmVtwoRequest() *TaobaoTrainAgentReturnticketConfirmVtwoRequest{
    return &TaobaoTrainAgentReturnticketConfirmVtwoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.returnticket.confirm.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerId Setter
// 用户id
func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetBuyerId(_buyerId int64) error {
    r._buyerId = _buyerId
    r.Set("buyer_id", _buyerId)
    return nil
}

// BuyerId Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetBuyerId() int64 {
    return r._buyerId
}
// AgreeReturn Setter
// 是否同意退票
func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetAgreeReturn(_agreeReturn bool) error {
    r._agreeReturn = _agreeReturn
    r.Set("agree_return", _agreeReturn)
    return nil
}

// AgreeReturn Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetAgreeReturn() bool {
    return r._agreeReturn
}
// RefundFee Setter
// 退票金额，单位分
func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetRefundFee(_refundFee int64) error {
    r._refundFee = _refundFee
    r.Set("refund_fee", _refundFee)
    return nil
}

// RefundFee Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetRefundFee() int64 {
    return r._refundFee
}
// MainBizOrderId Setter
// 淘宝主订单id
func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetMainBizOrderId(_mainBizOrderId int64) error {
    r._mainBizOrderId = _mainBizOrderId
    r.Set("main_biz_order_id", _mainBizOrderId)
    return nil
}

// MainBizOrderId Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetMainBizOrderId() int64 {
    return r._mainBizOrderId
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetAgentId() int64 {
    return r._agentId
}
// RefuseReturnReason Setter
// 拒绝退票原因
func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetRefuseReturnReason(_refuseReturnReason string) error {
    r._refuseReturnReason = _refuseReturnReason
    r.Set("refuse_return_reason", _refuseReturnReason)
    return nil
}

// RefuseReturnReason Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetRefuseReturnReason() string {
    return r._refuseReturnReason
}
// SubBizOrderId Setter
// 淘宝子订单id
func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetSubBizOrderId(_subBizOrderId int64) error {
    r._subBizOrderId = _subBizOrderId
    r.Set("sub_biz_order_id", _subBizOrderId)
    return nil
}

// SubBizOrderId Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetSubBizOrderId() int64 {
    return r._subBizOrderId
}
// CloseRefundNotify Setter
// 关闭通知用户 true:关闭 false:不关闭
func (r *TaobaoTrainAgentReturnticketConfirmVtwoRequest) SetCloseRefundNotify(_closeRefundNotify bool) error {
    r._closeRefundNotify = _closeRefundNotify
    r.Set("close_refund_notify", _closeRefundNotify)
    return nil
}

// CloseRefundNotify Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoRequest) GetCloseRefundNotify() bool {
    return r._closeRefundNotify
}
