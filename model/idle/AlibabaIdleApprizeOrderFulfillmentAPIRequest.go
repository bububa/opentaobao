package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleapprizeorderfulfillmentAPIRequest 鉴定担保资金订单履约 API请求
// alibaba.idle.apprize.order.fulfillment
//
// 服务商针对自己的服务订单进行履约
type AlibabaidleapprizeorderfulfillmentAPIRequest struct {
	model.Params
	// deal：服务商收取费用、refund：退款给付款方
	_action string
	// 天猫服务工单Id
	_workCardId int64
}

// NewAlibabaidleapprizeorderfulfillmentRequest 初始化AlibabaidleapprizeorderfulfillmentAPIRequest对象
func NewAlibabaidleapprizeorderfulfillmentRequest() *AlibabaidleapprizeorderfulfillmentAPIRequest {
	return &AlibabaidleapprizeorderfulfillmentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleapprizeorderfulfillmentAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.apprize.order.fulfillment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleapprizeorderfulfillmentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleapprizeorderfulfillmentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAction is Action Setter
// deal：服务商收取费用、refund：退款给付款方
func (r *AlibabaidleapprizeorderfulfillmentAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r AlibabaidleapprizeorderfulfillmentAPIRequest) GetAction() string {
	return r._action
}

// SetWorkCardId is WorkCardId Setter
// 天猫服务工单Id
func (r *AlibabaidleapprizeorderfulfillmentAPIRequest) SetWorkCardId(_workCardId int64) error {
	r._workCardId = _workCardId
	r.Set("work_card_id", _workCardId)
	return nil
}

// GetWorkCardId WorkCardId Getter
func (r AlibabaidleapprizeorderfulfillmentAPIRequest) GetWorkCardId() int64 {
	return r._workCardId
}
