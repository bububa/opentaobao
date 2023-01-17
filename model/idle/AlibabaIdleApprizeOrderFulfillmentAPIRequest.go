package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleApprizeOrderFulfillmentAPIRequest 鉴定担保资金订单履约 API请求
// alibaba.idle.apprize.order.fulfillment
//
// 服务商针对自己的服务订单进行履约
type AlibabaIdleApprizeOrderFulfillmentAPIRequest struct {
	model.Params
	// deal：服务商收取费用、refund：退款给付款方
	_action string
	// 天猫服务工单Id
	_workCardId int64
}

// NewAlibabaIdleApprizeOrderFulfillmentRequest 初始化AlibabaIdleApprizeOrderFulfillmentAPIRequest对象
func NewAlibabaIdleApprizeOrderFulfillmentRequest() *AlibabaIdleApprizeOrderFulfillmentAPIRequest {
	return &AlibabaIdleApprizeOrderFulfillmentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleApprizeOrderFulfillmentAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.apprize.order.fulfillment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleApprizeOrderFulfillmentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleApprizeOrderFulfillmentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAction is Action Setter
// deal：服务商收取费用、refund：退款给付款方
func (r *AlibabaIdleApprizeOrderFulfillmentAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r AlibabaIdleApprizeOrderFulfillmentAPIRequest) GetAction() string {
	return r._action
}

// SetWorkCardId is WorkCardId Setter
// 天猫服务工单Id
func (r *AlibabaIdleApprizeOrderFulfillmentAPIRequest) SetWorkCardId(_workCardId int64) error {
	r._workCardId = _workCardId
	r.Set("work_card_id", _workCardId)
	return nil
}

// GetWorkCardId WorkCardId Getter
func (r AlibabaIdleApprizeOrderFulfillmentAPIRequest) GetWorkCardId() int64 {
	return r._workCardId
}
