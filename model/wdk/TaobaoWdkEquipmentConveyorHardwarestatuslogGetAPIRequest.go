package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest 硬件状态变化日志查询 API请求
// taobao.wdk.equipment.conveyor.hardwarestatuslog.get
//
// 硬件状态变化日志查询
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) Reset() {
	r._warehouseId = 0
	r._conveyorId = 0
	r._startId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.hardwarestatuslog.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseId is WarehouseId Setter
// 仓库Id
func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) SetWarehouseId(_warehouseId int64) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) GetWarehouseId() int64 {
	return r._warehouseId
}

// SetConveyorId is ConveyorId Setter
// 悬挂链Id，即wcsNum
func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) SetConveyorId(_conveyorId int64) error {
	r._conveyorId = _conveyorId
	r.Set("conveyor_id", _conveyorId)
	return nil
}

// GetConveyorId ConveyorId Getter
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) GetConveyorId() int64 {
	return r._conveyorId
}

// SetStartId is StartId Setter
// 数据库id最小值
func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) SetStartId(_startId int64) error {
	r._startId = _startId
	r.Set("start_id", _startId)
	return nil
}

// GetStartId StartId Getter
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) GetStartId() int64 {
	return r._startId
}

var poolTaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest()
	},
}

// GetTaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest
func GetTaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest() *TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest {
	return poolTaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest.Get().(*TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest)
}

// ReleaseTaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest 将 TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest(v *TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest.Put(v)
}
