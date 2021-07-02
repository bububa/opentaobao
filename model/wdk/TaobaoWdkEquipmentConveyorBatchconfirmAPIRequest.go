package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest 五道口悬挂链信息批量确认 API请求
// taobao.wdk.equipment.conveyor.batchconfirm
//
// 批量消息确认
type TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest struct {
	model.Params
	// 仓库code
	_warehouseCode string
	// 待确认的uuid列表
	_uuids []string
}

// NewTaobaoWdkEquipmentConveyorBatchconfirmRequest 初始化TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest对象
func NewTaobaoWdkEquipmentConveyorBatchconfirmRequest() *TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest {
	return &TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.batchconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WarehouseCode Setter
// 仓库code
func (r *TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// Get WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// Set is Uuids Setter
// 待确认的uuid列表
func (r *TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) SetUuids(_uuids []string) error {
	r._uuids = _uuids
	r.Set("uuids", _uuids)
	return nil
}

// Get Uuids Getter
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetUuids() []string {
	return r._uuids
}
