package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalLogisticOrderCreateAPIRequest 创建物流订单 API请求
// cainiao.global.logistic.order.create
//
// 创建物流订单
type CainiaoGlobalLogisticOrderCreateAPIRequest struct {
	model.Params
	// 多语言
	_locale string
	// 订单参数
	_orderParam *OpenOrderParam
}

// NewCainiaoGlobalLogisticOrderCreateRequest 初始化CainiaoGlobalLogisticOrderCreateAPIRequest对象
func NewCainiaoGlobalLogisticOrderCreateRequest() *CainiaoGlobalLogisticOrderCreateAPIRequest {
	return &CainiaoGlobalLogisticOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalLogisticOrderCreateAPIRequest) GetApiMethodName() string {
	return "cainiao.global.logistic.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalLogisticOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoGlobalLogisticOrderCreateAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoGlobalLogisticOrderCreateAPIRequest) GetLocale() string {
	return r._locale
}

// SetOrderParam is OrderParam Setter
// 订单参数
func (r *CainiaoGlobalLogisticOrderCreateAPIRequest) SetOrderParam(_orderParam *OpenOrderParam) error {
	r._orderParam = _orderParam
	r.Set("order_param", _orderParam)
	return nil
}

// GetOrderParam OrderParam Getter
func (r CainiaoGlobalLogisticOrderCreateAPIRequest) GetOrderParam() *OpenOrderParam {
	return r._orderParam
}
