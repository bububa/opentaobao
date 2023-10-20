package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptraveltradetemplatequeryAPIRequest 订单服务详情模版查询 API请求
// alitrip.travel.trade.template.query
//
// 通过订单ID获取标注模版信息，商家可以根据模版来填充行业信息
type AlitriptraveltradetemplatequeryAPIRequest struct {
	model.Params
	// 淘宝平台订单ID
	_orderId int64
	// 是否取最新的模版
	_isNew bool
}

// NewAlitriptraveltradetemplatequeryRequest 初始化AlitriptraveltradetemplatequeryAPIRequest对象
func NewAlitriptraveltradetemplatequeryRequest() *AlitriptraveltradetemplatequeryAPIRequest {
	return &AlitriptraveltradetemplatequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptraveltradetemplatequeryAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.trade.template.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptraveltradetemplatequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptraveltradetemplatequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 淘宝平台订单ID
func (r *AlitriptraveltradetemplatequeryAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitriptraveltradetemplatequeryAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetIsNew is IsNew Setter
// 是否取最新的模版
func (r *AlitriptraveltradetemplatequeryAPIRequest) SetIsNew(_isNew bool) error {
	r._isNew = _isNew
	r.Set("is_new", _isNew)
	return nil
}

// GetIsNew IsNew Getter
func (r AlitriptraveltradetemplatequeryAPIRequest) GetIsNew() bool {
	return r._isNew
}
