package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophywarehouseordergetAPIRequest 仓作业单获取 API请求
// alibaba.tcls.aelophy.warehouse.order.get
//
// 仓作业单获取
type AlibabatclsaelophywarehouseordergetAPIRequest struct {
	model.Params
	// 查询入参对象
	_warehouseOrderGetRequest *WarehouseOrderGetRequest
}

// NewAlibabatclsaelophywarehouseordergetRequest 初始化AlibabatclsaelophywarehouseordergetAPIRequest对象
func NewAlibabatclsaelophywarehouseordergetRequest() *AlibabatclsaelophywarehouseordergetAPIRequest {
	return &AlibabatclsaelophywarehouseordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophywarehouseordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.warehouse.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophywarehouseordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophywarehouseordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseOrderGetRequest is WarehouseOrderGetRequest Setter
// 查询入参对象
func (r *AlibabatclsaelophywarehouseordergetAPIRequest) SetWarehouseOrderGetRequest(_warehouseOrderGetRequest *WarehouseOrderGetRequest) error {
	r._warehouseOrderGetRequest = _warehouseOrderGetRequest
	r.Set("warehouse_order_get_request", _warehouseOrderGetRequest)
	return nil
}

// GetWarehouseOrderGetRequest WarehouseOrderGetRequest Getter
func (r AlibabatclsaelophywarehouseordergetAPIRequest) GetWarehouseOrderGetRequest() *WarehouseOrderGetRequest {
	return r._warehouseOrderGetRequest
}
