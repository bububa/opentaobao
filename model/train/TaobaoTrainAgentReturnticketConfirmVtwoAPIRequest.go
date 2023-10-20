package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentreturnticketconfirmvtwoAPIRequest 退票通知 API请求
// taobao.train.agent.returnticket.confirm.vtwo
//
// 火车票代理商接口——退票通知回调
type TaobaotrainagentreturnticketconfirmvtwoAPIRequest struct {
	model.Params
	// 拒绝退票原因
	_refuseReturnReason string
	// 用户id
	_buyerId int64
	// 退票金额，单位分
	_refundFee int64
	// 淘宝主订单id
	_mainBizOrderId int64
	// 代理商id
	_agentId int64
	// 淘宝子订单id
	_subBizOrderId int64
	// 是否同意退票
	_agreeReturn bool
	// 关闭通知用户 true:关闭 false:不关闭
	_closeRefundNotify bool
}

// NewTaobaotrainagentreturnticketconfirmvtwoRequest 初始化TaobaotrainagentreturnticketconfirmvtwoAPIRequest对象
func NewTaobaotrainagentreturnticketconfirmvtwoRequest() *TaobaotrainagentreturnticketconfirmvtwoAPIRequest {
	return &TaobaotrainagentreturnticketconfirmvtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentreturnticketconfirmvtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.returnticket.confirm.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentreturnticketconfirmvtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentreturnticketconfirmvtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefuseReturnReason is RefuseReturnReason Setter
// 拒绝退票原因
func (r *TaobaotrainagentreturnticketconfirmvtwoAPIRequest) SetRefuseReturnReason(_refuseReturnReason string) error {
	r._refuseReturnReason = _refuseReturnReason
	r.Set("refuse_return_reason", _refuseReturnReason)
	return nil
}

// GetRefuseReturnReason RefuseReturnReason Getter
func (r TaobaotrainagentreturnticketconfirmvtwoAPIRequest) GetRefuseReturnReason() string {
	return r._refuseReturnReason
}

// SetBuyerId is BuyerId Setter
// 用户id
func (r *TaobaotrainagentreturnticketconfirmvtwoAPIRequest) SetBuyerId(_buyerId int64) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// GetBuyerId BuyerId Getter
func (r TaobaotrainagentreturnticketconfirmvtwoAPIRequest) GetBuyerId() int64 {
	return r._buyerId
}

// SetRefundFee is RefundFee Setter
// 退票金额，单位分
func (r *TaobaotrainagentreturnticketconfirmvtwoAPIRequest) SetRefundFee(_refundFee int64) error {
	r._refundFee = _refundFee
	r.Set("refund_fee", _refundFee)
	return nil
}

// GetRefundFee RefundFee Getter
func (r TaobaotrainagentreturnticketconfirmvtwoAPIRequest) GetRefundFee() int64 {
	return r._refundFee
}

// SetMainBizOrderId is MainBizOrderId Setter
// 淘宝主订单id
func (r *TaobaotrainagentreturnticketconfirmvtwoAPIRequest) SetMainBizOrderId(_mainBizOrderId int64) error {
	r._mainBizOrderId = _mainBizOrderId
	r.Set("main_biz_order_id", _mainBizOrderId)
	return nil
}

// GetMainBizOrderId MainBizOrderId Getter
func (r TaobaotrainagentreturnticketconfirmvtwoAPIRequest) GetMainBizOrderId() int64 {
	return r._mainBizOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaotrainagentreturnticketconfirmvtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaotrainagentreturnticketconfirmvtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetSubBizOrderId is SubBizOrderId Setter
// 淘宝子订单id
func (r *TaobaotrainagentreturnticketconfirmvtwoAPIRequest) SetSubBizOrderId(_subBizOrderId int64) error {
	r._subBizOrderId = _subBizOrderId
	r.Set("sub_biz_order_id", _subBizOrderId)
	return nil
}

// GetSubBizOrderId SubBizOrderId Getter
func (r TaobaotrainagentreturnticketconfirmvtwoAPIRequest) GetSubBizOrderId() int64 {
	return r._subBizOrderId
}

// SetAgreeReturn is AgreeReturn Setter
// 是否同意退票
func (r *TaobaotrainagentreturnticketconfirmvtwoAPIRequest) SetAgreeReturn(_agreeReturn bool) error {
	r._agreeReturn = _agreeReturn
	r.Set("agree_return", _agreeReturn)
	return nil
}

// GetAgreeReturn AgreeReturn Getter
func (r TaobaotrainagentreturnticketconfirmvtwoAPIRequest) GetAgreeReturn() bool {
	return r._agreeReturn
}

// SetCloseRefundNotify is CloseRefundNotify Setter
// 关闭通知用户 true:关闭 false:不关闭
func (r *TaobaotrainagentreturnticketconfirmvtwoAPIRequest) SetCloseRefundNotify(_closeRefundNotify bool) error {
	r._closeRefundNotify = _closeRefundNotify
	r.Set("close_refund_notify", _closeRefundNotify)
	return nil
}

// GetCloseRefundNotify CloseRefundNotify Getter
func (r TaobaotrainagentreturnticketconfirmvtwoAPIRequest) GetCloseRefundNotify() bool {
	return r._closeRefundNotify
}
