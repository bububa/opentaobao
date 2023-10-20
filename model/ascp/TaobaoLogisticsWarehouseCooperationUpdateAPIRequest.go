package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswarehousecooperationupdateAPIRequest 合作商家信息同步 API请求
// taobao.logistics.warehouse.cooperation.update
//
// 合作商家信息同步
type TaobaologisticswarehousecooperationupdateAPIRequest struct {
	model.Params
	// 合作商家信息同步入参
	_warehouseCooperationUpdateRequest *WarehouseCooperationUpdateRequest
}

// NewTaobaologisticswarehousecooperationupdateRequest 初始化TaobaologisticswarehousecooperationupdateAPIRequest对象
func NewTaobaologisticswarehousecooperationupdateRequest() *TaobaologisticswarehousecooperationupdateAPIRequest {
	return &TaobaologisticswarehousecooperationupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticswarehousecooperationupdateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.warehouse.cooperation.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticswarehousecooperationupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticswarehousecooperationupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCooperationUpdateRequest is WarehouseCooperationUpdateRequest Setter
// 合作商家信息同步入参
func (r *TaobaologisticswarehousecooperationupdateAPIRequest) SetWarehouseCooperationUpdateRequest(_warehouseCooperationUpdateRequest *WarehouseCooperationUpdateRequest) error {
	r._warehouseCooperationUpdateRequest = _warehouseCooperationUpdateRequest
	r.Set("warehouse_cooperation_update_request", _warehouseCooperationUpdateRequest)
	return nil
}

// GetWarehouseCooperationUpdateRequest WarehouseCooperationUpdateRequest Getter
func (r TaobaologisticswarehousecooperationupdateAPIRequest) GetWarehouseCooperationUpdateRequest() *WarehouseCooperationUpdateRequest {
	return r._warehouseCooperationUpdateRequest
}
