package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest
硬件状态变化日志查询 API请求
taobao.wdk.equipment.conveyor.hardwarestatuslog.get

硬件状态变化日志查询 */
type TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest struct {
	model.Params
	// 仓库Id
	_warehouseId int64
	// 悬挂链Id，即wcsNum
	_conveyorId int64
	// 数据库id最小值
	_startId int64
}

// NewTaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest 初始化TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest对象
func NewTaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest() *TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest {
	return &TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.hardwarestatuslog.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WarehouseId Setter
// 仓库Id
func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) SetWarehouseId(_warehouseId int64) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// Get WarehouseId Getter
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) GetWarehouseId() int64 {
	return r._warehouseId
}

// Set is ConveyorId Setter
// 悬挂链Id，即wcsNum
func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) SetConveyorId(_conveyorId int64) error {
	r._conveyorId = _conveyorId
	r.Set("conveyor_id", _conveyorId)
	return nil
}

// Get ConveyorId Getter
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) GetConveyorId() int64 {
	return r._conveyorId
}

// Set is StartId Setter
// 数据库id最小值
func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) SetStartId(_startId int64) error {
	r._startId = _startId
	r.Set("start_id", _startId)
	return nil
}

// Get StartId Getter
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) GetStartId() int64 {
	return r._startId
}
