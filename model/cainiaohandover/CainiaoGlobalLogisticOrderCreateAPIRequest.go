package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaogloballogisticordercreateAPIRequest 创建物流订单 API请求
// cainiao.global.logistic.order.create
//
// 创建物流订单
type CainiaogloballogisticordercreateAPIRequest struct {
	model.Params
	// 多语言
	_locale string
	// 订单参数
	_orderParam *OpenOrderParam
}

// NewCainiaogloballogisticordercreateRequest 初始化CainiaogloballogisticordercreateAPIRequest对象
func NewCainiaogloballogisticordercreateRequest() *CainiaogloballogisticordercreateAPIRequest {
	return &CainiaogloballogisticordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaogloballogisticordercreateAPIRequest) GetApiMethodName() string {
	return "cainiao.global.logistic.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaogloballogisticordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaogloballogisticordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaogloballogisticordercreateAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaogloballogisticordercreateAPIRequest) GetLocale() string {
	return r._locale
}

// SetOrderParam is OrderParam Setter
// 订单参数
func (r *CainiaogloballogisticordercreateAPIRequest) SetOrderParam(_orderParam *OpenOrderParam) error {
	r._orderParam = _orderParam
	r.Set("order_param", _orderParam)
	return nil
}

// GetOrderParam OrderParam Getter
func (r CainiaogloballogisticordercreateAPIRequest) GetOrderParam() *OpenOrderParam {
	return r._orderParam
}
