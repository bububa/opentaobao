package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswarehouseoperationupdateAPIRequest 仓作业能力新建/更新 API请求
// taobao.logistics.warehouse.operation.update
//
// 仓作业能力新建/更新
type TaobaologisticswarehouseoperationupdateAPIRequest struct {
	model.Params
	// 仓作业能力新建/更新入参
	_warehouseOperationUpdateRequest *WarehouseOperationUpdateRequest
}

// NewTaobaologisticswarehouseoperationupdateRequest 初始化TaobaologisticswarehouseoperationupdateAPIRequest对象
func NewTaobaologisticswarehouseoperationupdateRequest() *TaobaologisticswarehouseoperationupdateAPIRequest {
	return &TaobaologisticswarehouseoperationupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticswarehouseoperationupdateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.warehouse.operation.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticswarehouseoperationupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticswarehouseoperationupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseOperationUpdateRequest is WarehouseOperationUpdateRequest Setter
// 仓作业能力新建/更新入参
func (r *TaobaologisticswarehouseoperationupdateAPIRequest) SetWarehouseOperationUpdateRequest(_warehouseOperationUpdateRequest *WarehouseOperationUpdateRequest) error {
	r._warehouseOperationUpdateRequest = _warehouseOperationUpdateRequest
	r.Set("warehouse_operation_update_request", _warehouseOperationUpdateRequest)
	return nil
}

// GetWarehouseOperationUpdateRequest WarehouseOperationUpdateRequest Getter
func (r TaobaologisticswarehouseoperationupdateAPIRequest) GetWarehouseOperationUpdateRequest() *WarehouseOperationUpdateRequest {
	return r._warehouseOperationUpdateRequest
}
