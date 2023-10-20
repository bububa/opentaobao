package healthnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthnrlogisticsqueryAPIRequest 阿里健康新零售物流详情接口 API请求
// alibaba.health.nr.logistics.query
//
// 对阿里健康o2o对接的商户提供查询物流单详情的能力
type AlibabahealthnrlogisticsqueryAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewAlibabahealthnrlogisticsqueryRequest 初始化AlibabahealthnrlogisticsqueryAPIRequest对象
func NewAlibabahealthnrlogisticsqueryRequest() *AlibabahealthnrlogisticsqueryAPIRequest {
	return &AlibabahealthnrlogisticsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthnrlogisticsqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.health.nr.logistics.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthnrlogisticsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthnrlogisticsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabahealthnrlogisticsqueryAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahealthnrlogisticsqueryAPIRequest) GetOrderId() int64 {
	return r._orderId
}
