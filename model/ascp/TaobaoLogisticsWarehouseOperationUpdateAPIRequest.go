package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWarehouseOperationUpdateAPIRequest 仓作业能力新建/更新 API请求
// taobao.logistics.warehouse.operation.update
//
// 仓作业能力新建/更新
type TaobaoLogisticsWarehouseOperationUpdateAPIRequest struct {
	model.Params
	// 仓作业能力新建/更新入参
	_warehouseOperationUpdateRequest *WarehouseOperationUpdateRequest
}

// NewTaobaoLogisticsWarehouseOperationUpdateRequest 初始化TaobaoLogisticsWarehouseOperationUpdateAPIRequest对象
func NewTaobaoLogisticsWarehouseOperationUpdateRequest() *TaobaoLogisticsWarehouseOperationUpdateAPIRequest {
	return &TaobaoLogisticsWarehouseOperationUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWarehouseOperationUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.warehouse.operation.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWarehouseOperationUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWarehouseOperationUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseOperationUpdateRequest is WarehouseOperationUpdateRequest Setter
// 仓作业能力新建/更新入参
func (r *TaobaoLogisticsWarehouseOperationUpdateAPIRequest) SetWarehouseOperationUpdateRequest(_warehouseOperationUpdateRequest *WarehouseOperationUpdateRequest) error {
	r._warehouseOperationUpdateRequest = _warehouseOperationUpdateRequest
	r.Set("warehouse_operation_update_request", _warehouseOperationUpdateRequest)
	return nil
}

// GetWarehouseOperationUpdateRequest WarehouseOperationUpdateRequest Getter
func (r TaobaoLogisticsWarehouseOperationUpdateAPIRequest) GetWarehouseOperationUpdateRequest() *WarehouseOperationUpdateRequest {
	return r._warehouseOperationUpdateRequest
}
