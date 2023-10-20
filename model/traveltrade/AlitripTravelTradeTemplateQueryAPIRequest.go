package traveltrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelTradeTemplateQueryAPIRequest 订单服务详情模版查询 API请求
// alitrip.travel.trade.template.query
//
// 通过订单ID获取标注模版信息，商家可以根据模版来填充行业信息
type AlitripTravelTradeTemplateQueryAPIRequest struct {
	model.Params
	// 淘宝平台订单ID
	_orderId int64
	// 是否取最新的模版
	_isNew bool
}

// NewAlitripTravelTradeTemplateQueryRequest 初始化AlitripTravelTradeTemplateQueryAPIRequest对象
func NewAlitripTravelTradeTemplateQueryRequest() *AlitripTravelTradeTemplateQueryAPIRequest {
	return &AlitripTravelTradeTemplateQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelTradeTemplateQueryAPIRequest) Reset() {
	r._orderId = 0
	r._isNew = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelTradeTemplateQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.trade.template.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelTradeTemplateQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelTradeTemplateQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripTravelTradeTemplateQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelTradeTemplateQueryRequest()
	},
}

// GetAlitripTravelTradeTemplateQueryRequest 从 sync.Pool 获取 AlitripTravelTradeTemplateQueryAPIRequest
func GetAlitripTravelTradeTemplateQueryAPIRequest() *AlitripTravelTradeTemplateQueryAPIRequest {
	return poolAlitripTravelTradeTemplateQueryAPIRequest.Get().(*AlitripTravelTradeTemplateQueryAPIRequest)
}

// ReleaseAlitripTravelTradeTemplateQueryAPIRequest 将 AlitripTravelTradeTemplateQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelTradeTemplateQueryAPIRequest(v *AlitripTravelTradeTemplateQueryAPIRequest) {
	v.Reset()
	poolAlitripTravelTradeTemplateQueryAPIRequest.Put(v)
}
