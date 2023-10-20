package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentconveyorconveyorinfogetAPIRequest 获取五道口悬挂链信息 API请求
// taobao.wdk.equipment.conveyor.conveyorinfo.get
//
// 获取五道口悬挂链信息
type TaobaowdkequipmentconveyorconveyorinfogetAPIRequest struct {
	model.Params
	// 仓库code
	_warehouseCode string
	// wcsNum
	_conveyorId int64
}

// NewTaobaowdkequipmentconveyorconveyorinfogetRequest 初始化TaobaowdkequipmentconveyorconveyorinfogetAPIRequest对象
func NewTaobaowdkequipmentconveyorconveyorinfogetRequest() *TaobaowdkequipmentconveyorconveyorinfogetAPIRequest {
	return &TaobaowdkequipmentconveyorconveyorinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowdkequipmentconveyorconveyorinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.conveyorinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowdkequipmentconveyorconveyorinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowdkequipmentconveyorconveyorinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 仓库code
func (r *TaobaowdkequipmentconveyorconveyorinfogetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaowdkequipmentconveyorconveyorinfogetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetConveyorId is ConveyorId Setter
// wcsNum
func (r *TaobaowdkequipmentconveyorconveyorinfogetAPIRequest) SetConveyorId(_conveyorId int64) error {
	r._conveyorId = _conveyorId
	r.Set("conveyor_id", _conveyorId)
	return nil
}

// GetConveyorId ConveyorId Getter
func (r TaobaowdkequipmentconveyorconveyorinfogetAPIRequest) GetConveyorId() int64 {
	return r._conveyorId
}
