package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest 退票通知 API请求
// taobao.train.agent.returnticket.confirm.vtwo
//
// 火车票代理商接口——退票通知回调
type TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest struct {
	model.Params
	// 用户id
	_buyerId int64
	// 是否同意退票
	_agreeReturn bool
	// 退票金额，单位分
	_refundFee int64
	// 淘宝主订单id
	_mainBizOrderId int64
	// 代理商id
	_agentId int64
	// 拒绝退票原因
	_refuseReturnReason string
	// 淘宝子订单id
	_subBizOrderId int64
	// 关闭通知用户 true:关闭 false:不关闭
	_closeRefundNotify bool
}

// NewTaobaoTrainAgentReturnticketConfirmVtwoRequest 初始化TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest对象
func NewTaobaoTrainAgentReturnticketConfirmVtwoRequest() *TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest {
	return &TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.returnticket.confirm.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BuyerId Setter
// 用户id
func (r *TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) SetBuyerId(_buyerId int64) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// Get BuyerId Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) GetBuyerId() int64 {
	return r._buyerId
}

// Set is AgreeReturn Setter
// 是否同意退票
func (r *TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) SetAgreeReturn(_agreeReturn bool) error {
	r._agreeReturn = _agreeReturn
	r.Set("agree_return", _agreeReturn)
	return nil
}

// Get AgreeReturn Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) GetAgreeReturn() bool {
	return r._agreeReturn
}

// Set is RefundFee Setter
// 退票金额，单位分
func (r *TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) SetRefundFee(_refundFee int64) error {
	r._refundFee = _refundFee
	r.Set("refund_fee", _refundFee)
	return nil
}

// Get RefundFee Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) GetRefundFee() int64 {
	return r._refundFee
}

// Set is MainBizOrderId Setter
// 淘宝主订单id
func (r *TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) SetMainBizOrderId(_mainBizOrderId int64) error {
	r._mainBizOrderId = _mainBizOrderId
	r.Set("main_biz_order_id", _mainBizOrderId)
	return nil
}

// Get MainBizOrderId Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) GetMainBizOrderId() int64 {
	return r._mainBizOrderId
}

// Set is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// Get AgentId Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// Set is RefuseReturnReason Setter
// 拒绝退票原因
func (r *TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) SetRefuseReturnReason(_refuseReturnReason string) error {
	r._refuseReturnReason = _refuseReturnReason
	r.Set("refuse_return_reason", _refuseReturnReason)
	return nil
}

// Get RefuseReturnReason Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) GetRefuseReturnReason() string {
	return r._refuseReturnReason
}

// Set is SubBizOrderId Setter
// 淘宝子订单id
func (r *TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) SetSubBizOrderId(_subBizOrderId int64) error {
	r._subBizOrderId = _subBizOrderId
	r.Set("sub_biz_order_id", _subBizOrderId)
	return nil
}

// Get SubBizOrderId Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) GetSubBizOrderId() int64 {
	return r._subBizOrderId
}

// Set is CloseRefundNotify Setter
// 关闭通知用户 true:关闭 false:不关闭
func (r *TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) SetCloseRefundNotify(_closeRefundNotify bool) error {
	r._closeRefundNotify = _closeRefundNotify
	r.Set("close_refund_notify", _closeRefundNotify)
	return nil
}

// Get CloseRefundNotify Getter
func (r TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest) GetCloseRefundNotify() bool {
	return r._closeRefundNotify
}
