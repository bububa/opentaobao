package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptraveltradequeryAPIRequest 飞猪度假-订单详情查询接口 API请求
// alitrip.travel.trade.query
//
// 飞猪度假订单详情查询接口
type AlitriptraveltradequeryAPIRequest struct {
	model.Params
	// 主订单id
	_orderId int64
}

// NewAlitriptraveltradequeryRequest 初始化AlitriptraveltradequeryAPIRequest对象
func NewAlitriptraveltradequeryRequest() *AlitriptraveltradequeryAPIRequest {
	return &AlitriptraveltradequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptraveltradequeryAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.trade.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptraveltradequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptraveltradequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 主订单id
func (r *AlitriptraveltradequeryAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitriptraveltradequeryAPIRequest) GetOrderId() int64 {
	return r._orderId
}
