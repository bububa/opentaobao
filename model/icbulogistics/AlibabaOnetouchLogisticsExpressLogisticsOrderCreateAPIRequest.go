package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpresslogisticsordercreateAPIRequest 快递下单 API请求
// alibaba.onetouch.logistics.express.logistics.order.create
//
// 快递下单
type AlibabaonetouchlogisticsexpresslogisticsordercreateAPIRequest struct {
	model.Params
	// 请求参数对象
	_paramnQuery *PlaceOrderDto
}

// NewAlibabaonetouchlogisticsexpresslogisticsordercreateRequest 初始化AlibabaonetouchlogisticsexpresslogisticsordercreateAPIRequest对象
func NewAlibabaonetouchlogisticsexpresslogisticsordercreateRequest() *AlibabaonetouchlogisticsexpresslogisticsordercreateAPIRequest {
	return &AlibabaonetouchlogisticsexpresslogisticsordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaonetouchlogisticsexpresslogisticsordercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.logistics.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaonetouchlogisticsexpresslogisticsordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaonetouchlogisticsexpresslogisticsordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamnQuery is ParamnQuery Setter
// 请求参数对象
func (r *AlibabaonetouchlogisticsexpresslogisticsordercreateAPIRequest) SetParamnQuery(_paramnQuery *PlaceOrderDto) error {
	r._paramnQuery = _paramnQuery
	r.Set("paramn_query", _paramnQuery)
	return nil
}

// GetParamnQuery ParamnQuery Getter
func (r AlibabaonetouchlogisticsexpresslogisticsordercreateAPIRequest) GetParamnQuery() *PlaceOrderDto {
	return r._paramnQuery
}
