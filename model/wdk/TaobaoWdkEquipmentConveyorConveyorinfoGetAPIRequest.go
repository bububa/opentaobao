package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest 获取五道口悬挂链信息 API请求
// taobao.wdk.equipment.conveyor.conveyorinfo.get
//
// 获取五道口悬挂链信息
type TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest struct {
	model.Params
	// 仓库code
	_warehouseCode string
	// wcsNum
	_conveyorId int64
}

// NewTaobaoWdkEquipmentConveyorConveyorinfoGetRequest 初始化TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest对象
func NewTaobaoWdkEquipmentConveyorConveyorinfoGetRequest() *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest {
	return &TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.conveyorinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWarehouseCode is WarehouseCode Setter
// 仓库code
func (r *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetConveyorId is ConveyorId Setter
// wcsNum
func (r *TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) SetConveyorId(_conveyorId int64) error {
	r._conveyorId = _conveyorId
	r.Set("conveyor_id", _conveyorId)
	return nil
}

// GetConveyorId ConveyorId Getter
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest) GetConveyorId() int64 {
	return r._conveyorId
}
