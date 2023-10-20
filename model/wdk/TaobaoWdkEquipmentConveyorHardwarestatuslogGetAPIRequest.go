package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest 硬件状态变化日志查询 API请求
// taobao.wdk.equipment.conveyor.hardwarestatuslog.get
//
// 硬件状态变化日志查询
type TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest struct {
	model.Params
	// 仓库Id
	_warehouseId int64
	// 悬挂链Id，即wcsNum
	_conveyorId int64
	// 数据库id最小值
	_startId int64
}

// NewTaobaowdkequipmentconveyorhardwarestatusloggetRequest 初始化TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest对象
func NewTaobaowdkequipmentconveyorhardwarestatusloggetRequest() *TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest {
	return &TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.hardwarestatuslog.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseId is WarehouseId Setter
// 仓库Id
func (r *TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest) SetWarehouseId(_warehouseId int64) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest) GetWarehouseId() int64 {
	return r._warehouseId
}

// SetConveyorId is ConveyorId Setter
// 悬挂链Id，即wcsNum
func (r *TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest) SetConveyorId(_conveyorId int64) error {
	r._conveyorId = _conveyorId
	r.Set("conveyor_id", _conveyorId)
	return nil
}

// GetConveyorId ConveyorId Getter
func (r TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest) GetConveyorId() int64 {
	return r._conveyorId
}

// SetStartId is StartId Setter
// 数据库id最小值
func (r *TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest) SetStartId(_startId int64) error {
	r._startId = _startId
	r.Set("start_id", _startId)
	return nil
}

// GetStartId StartId Getter
func (r TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest) GetStartId() int64 {
	return r._startId
}
