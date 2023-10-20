package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWarehouseResourceUpdateAPIRequest 服务商新建/更新仓 API请求
// taobao.logistics.warehouse.resource.update
//
// 服务商新建/更新仓
type TaobaoLogisticsWarehouseResourceUpdateAPIRequest struct {
	model.Params
	// 创建/更新仓入参
	_warehouseUpdateRequest *WarehouseUpdateRequest
}

// NewTaobaoLogisticsWarehouseResourceUpdateRequest 初始化TaobaoLogisticsWarehouseResourceUpdateAPIRequest对象
func NewTaobaoLogisticsWarehouseResourceUpdateRequest() *TaobaoLogisticsWarehouseResourceUpdateAPIRequest {
	return &TaobaoLogisticsWarehouseResourceUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWarehouseResourceUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.warehouse.resource.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWarehouseResourceUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWarehouseResourceUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseUpdateRequest is WarehouseUpdateRequest Setter
// 创建/更新仓入参
func (r *TaobaoLogisticsWarehouseResourceUpdateAPIRequest) SetWarehouseUpdateRequest(_warehouseUpdateRequest *WarehouseUpdateRequest) error {
	r._warehouseUpdateRequest = _warehouseUpdateRequest
	r.Set("warehouse_update_request", _warehouseUpdateRequest)
	return nil
}

// GetWarehouseUpdateRequest WarehouseUpdateRequest Getter
func (r TaobaoLogisticsWarehouseResourceUpdateAPIRequest) GetWarehouseUpdateRequest() *WarehouseUpdateRequest {
	return r._warehouseUpdateRequest
}
