package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyWarehouseOrderGetAPIRequest 仓作业单获取 API请求
// alibaba.tcls.aelophy.warehouse.order.get
//
// 仓作业单获取
type AlibabaTclsAelophyWarehouseOrderGetAPIRequest struct {
	model.Params
	// 查询入参对象
	_warehouseOrderGetRequest *WarehouseOrderGetRequest
}

// NewAlibabaTclsAelophyWarehouseOrderGetRequest 初始化AlibabaTclsAelophyWarehouseOrderGetAPIRequest对象
func NewAlibabaTclsAelophyWarehouseOrderGetRequest() *AlibabaTclsAelophyWarehouseOrderGetAPIRequest {
	return &AlibabaTclsAelophyWarehouseOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyWarehouseOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.warehouse.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyWarehouseOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyWarehouseOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseOrderGetRequest is WarehouseOrderGetRequest Setter
// 查询入参对象
func (r *AlibabaTclsAelophyWarehouseOrderGetAPIRequest) SetWarehouseOrderGetRequest(_warehouseOrderGetRequest *WarehouseOrderGetRequest) error {
	r._warehouseOrderGetRequest = _warehouseOrderGetRequest
	r.Set("warehouse_order_get_request", _warehouseOrderGetRequest)
	return nil
}

// GetWarehouseOrderGetRequest WarehouseOrderGetRequest Getter
func (r AlibabaTclsAelophyWarehouseOrderGetAPIRequest) GetWarehouseOrderGetRequest() *WarehouseOrderGetRequest {
	return r._warehouseOrderGetRequest
}
