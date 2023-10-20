package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenSellerBizLogisticTimeRuleAPIRequest 商家自定义发货时效 API请求
// taobao.open.seller.biz.logistic.time.rule
//
// 服务商回传商家自定义发货时效
type TaobaoOpenSellerBizLogisticTimeRuleAPIRequest struct {
	model.Params
	// 当日订单最晚截单时间， 当日可发货订单最晚的支付时间 24小时制 格式：HH:mm
	_lastPayTime string
	// 当日仓库最晚出库时间 当日可发货订单最晚的出库时间 24小时制 格式：HH:mm
	_lastDeliveryTime string
}

// NewTaobaoOpenSellerBizLogisticTimeRuleRequest 初始化TaobaoOpenSellerBizLogisticTimeRuleAPIRequest对象
func NewTaobaoOpenSellerBizLogisticTimeRuleRequest() *TaobaoOpenSellerBizLogisticTimeRuleAPIRequest {
	return &TaobaoOpenSellerBizLogisticTimeRuleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenSellerBizLogisticTimeRuleAPIRequest) GetApiMethodName() string {
	return "taobao.open.seller.biz.logistic.time.rule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenSellerBizLogisticTimeRuleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenSellerBizLogisticTimeRuleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLastPayTime is LastPayTime Setter
// 当日订单最晚截单时间， 当日可发货订单最晚的支付时间 24小时制 格式：HH:mm
func (r *TaobaoOpenSellerBizLogisticTimeRuleAPIRequest) SetLastPayTime(_lastPayTime string) error {
	r._lastPayTime = _lastPayTime
	r.Set("last_pay_time", _lastPayTime)
	return nil
}

// GetLastPayTime LastPayTime Getter
func (r TaobaoOpenSellerBizLogisticTimeRuleAPIRequest) GetLastPayTime() string {
	return r._lastPayTime
}

// SetLastDeliveryTime is LastDeliveryTime Setter
// 当日仓库最晚出库时间 当日可发货订单最晚的出库时间 24小时制 格式：HH:mm
func (r *TaobaoOpenSellerBizLogisticTimeRuleAPIRequest) SetLastDeliveryTime(_lastDeliveryTime string) error {
	r._lastDeliveryTime = _lastDeliveryTime
	r.Set("last_delivery_time", _lastDeliveryTime)
	return nil
}

// GetLastDeliveryTime LastDeliveryTime Getter
func (r TaobaoOpenSellerBizLogisticTimeRuleAPIRequest) GetLastDeliveryTime() string {
	return r._lastDeliveryTime
}
