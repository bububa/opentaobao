package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest 异常通道日志查询 API请求
// taobao.wdk.equipment.conveyor.exceptionslidewaylog.get
//
// 五道口悬挂链异常通道事件查询
type TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest struct {
	model.Params
	// 仓库Id
	_warehouseId int64
	// 悬挂链Id，即wcsNum
	_conveyorId int64
	// 数据库id最小值
	_startId int64
}

// NewTaobaowdkequipmentconveyorexceptionslidewayloggetRequest 初始化TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest对象
func NewTaobaowdkequipmentconveyorexceptionslidewayloggetRequest() *TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest {
	return &TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.exceptionslidewaylog.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseId is WarehouseId Setter
// 仓库Id
func (r *TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest) SetWarehouseId(_warehouseId int64) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest) GetWarehouseId() int64 {
	return r._warehouseId
}

// SetConveyorId is ConveyorId Setter
// 悬挂链Id，即wcsNum
func (r *TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest) SetConveyorId(_conveyorId int64) error {
	r._conveyorId = _conveyorId
	r.Set("conveyor_id", _conveyorId)
	return nil
}

// GetConveyorId ConveyorId Getter
func (r TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest) GetConveyorId() int64 {
	return r._conveyorId
}

// SetStartId is StartId Setter
// 数据库id最小值
func (r *TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest) SetStartId(_startId int64) error {
	r._startId = _startId
	r.Set("start_id", _startId)
	return nil
}

// GetStartId StartId Getter
func (r TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest) GetStartId() int64 {
	return r._startId
}
