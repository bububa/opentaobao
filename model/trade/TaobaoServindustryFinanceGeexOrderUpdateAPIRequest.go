package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoServindustryFinanceGeexOrderUpdateAPIRequest 即科订单结果更新回调 API请求
// taobao.servindustry.finance.geex.order.update
//
// 即科订单结果更新回调本地接口
type TaobaoServindustryFinanceGeexOrderUpdateAPIRequest struct {
	model.Params
	// 原因
	_reason string
	// 订单号
	_alscOrderId string
	// 状态
	_applyStatus string
}

// NewTaobaoServindustryFinanceGeexOrderUpdateRequest 初始化TaobaoServindustryFinanceGeexOrderUpdateAPIRequest对象
func NewTaobaoServindustryFinanceGeexOrderUpdateRequest() *TaobaoServindustryFinanceGeexOrderUpdateAPIRequest {
	return &TaobaoServindustryFinanceGeexOrderUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoServindustryFinanceGeexOrderUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.servindustry.finance.geex.order.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoServindustryFinanceGeexOrderUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoServindustryFinanceGeexOrderUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReason is Reason Setter
// 原因
func (r *TaobaoServindustryFinanceGeexOrderUpdateAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoServindustryFinanceGeexOrderUpdateAPIRequest) GetReason() string {
	return r._reason
}

// SetAlscOrderId is AlscOrderId Setter
// 订单号
func (r *TaobaoServindustryFinanceGeexOrderUpdateAPIRequest) SetAlscOrderId(_alscOrderId string) error {
	r._alscOrderId = _alscOrderId
	r.Set("alsc_order_id", _alscOrderId)
	return nil
}

// GetAlscOrderId AlscOrderId Getter
func (r TaobaoServindustryFinanceGeexOrderUpdateAPIRequest) GetAlscOrderId() string {
	return r._alscOrderId
}

// SetApplyStatus is ApplyStatus Setter
// 状态
func (r *TaobaoServindustryFinanceGeexOrderUpdateAPIRequest) SetApplyStatus(_applyStatus string) error {
	r._applyStatus = _applyStatus
	r.Set("apply_status", _applyStatus)
	return nil
}

// GetApplyStatus ApplyStatus Getter
func (r TaobaoServindustryFinanceGeexOrderUpdateAPIRequest) GetApplyStatus() string {
	return r._applyStatus
}
