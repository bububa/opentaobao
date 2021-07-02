package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelTradeTemplateQueryAPIRequest 订单服务详情模版查询 API请求
// alitrip.travel.trade.template.query
//
// 通过订单ID获取标注模版信息，商家可以根据模版来填充行业信息
type AlitripTravelTradeTemplateQueryAPIRequest struct {
	model.Params
	// 是否取最新的模版
	_isNew bool
	// 淘宝平台订单ID
	_orderId int64
}

// NewAlitripTravelTradeTemplateQueryRequest 初始化AlitripTravelTradeTemplateQueryAPIRequest对象
func NewAlitripTravelTradeTemplateQueryRequest() *AlitripTravelTradeTemplateQueryAPIRequest {
	return &AlitripTravelTradeTemplateQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelTradeTemplateQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.trade.template.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelTradeTemplateQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetIsNew is IsNew Setter
// 是否取最新的模版
func (r *AlitripTravelTradeTemplateQueryAPIRequest) SetIsNew(_isNew bool) error {
	r._isNew = _isNew
	r.Set("is_new", _isNew)
	return nil
}

// GetIsNew IsNew Getter
func (r AlitripTravelTradeTemplateQueryAPIRequest) GetIsNew() bool {
	return r._isNew
}

// SetOrderId is OrderId Setter
// 淘宝平台订单ID
func (r *AlitripTravelTradeTemplateQueryAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripTravelTradeTemplateQueryAPIRequest) GetOrderId() int64 {
	return r._orderId
}
