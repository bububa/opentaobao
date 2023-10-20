package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswarehousecooperationbatchconfirmAPIRequest 仓合作关系确认 API请求
// taobao.logistics.warehouse.cooperation.batch.confirm
//
// 仓合作关系确认
type TaobaologisticswarehousecooperationbatchconfirmAPIRequest struct {
	model.Params
	// 仓合作关系确认入参
	_warehouseCooperationBatchConfirmRequest *WarehouseCooperationBatchConfirmRequest
}

// NewTaobaologisticswarehousecooperationbatchconfirmRequest 初始化TaobaologisticswarehousecooperationbatchconfirmAPIRequest对象
func NewTaobaologisticswarehousecooperationbatchconfirmRequest() *TaobaologisticswarehousecooperationbatchconfirmAPIRequest {
	return &TaobaologisticswarehousecooperationbatchconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticswarehousecooperationbatchconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.warehouse.cooperation.batch.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticswarehousecooperationbatchconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticswarehousecooperationbatchconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCooperationBatchConfirmRequest is WarehouseCooperationBatchConfirmRequest Setter
// 仓合作关系确认入参
func (r *TaobaologisticswarehousecooperationbatchconfirmAPIRequest) SetWarehouseCooperationBatchConfirmRequest(_warehouseCooperationBatchConfirmRequest *WarehouseCooperationBatchConfirmRequest) error {
	r._warehouseCooperationBatchConfirmRequest = _warehouseCooperationBatchConfirmRequest
	r.Set("warehouse_cooperation_batch_confirm_request", _warehouseCooperationBatchConfirmRequest)
	return nil
}

// GetWarehouseCooperationBatchConfirmRequest WarehouseCooperationBatchConfirmRequest Getter
func (r TaobaologisticswarehousecooperationbatchconfirmAPIRequest) GetWarehouseCooperationBatchConfirmRequest() *WarehouseCooperationBatchConfirmRequest {
	return r._warehouseCooperationBatchConfirmRequest
}
