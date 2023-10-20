package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentconveyorstatusconfirmAPIRequest 悬挂链状态回传确认 API请求
// taobao.wdk.equipment.conveyor.statusconfirm
//
// 悬挂链状态回传确认
type TaobaowdkequipmentconveyorstatusconfirmAPIRequest struct {
	model.Params
	// warehouse_code
	_warehouseCode string
	// uuid
	_uuid string
}

// NewTaobaowdkequipmentconveyorstatusconfirmRequest 初始化TaobaowdkequipmentconveyorstatusconfirmAPIRequest对象
func NewTaobaowdkequipmentconveyorstatusconfirmRequest() *TaobaowdkequipmentconveyorstatusconfirmAPIRequest {
	return &TaobaowdkequipmentconveyorstatusconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowdkequipmentconveyorstatusconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.statusconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowdkequipmentconveyorstatusconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowdkequipmentconveyorstatusconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// warehouse_code
func (r *TaobaowdkequipmentconveyorstatusconfirmAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaowdkequipmentconveyorstatusconfirmAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetUuid is Uuid Setter
// uuid
func (r *TaobaowdkequipmentconveyorstatusconfirmAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaowdkequipmentconveyorstatusconfirmAPIRequest) GetUuid() string {
	return r._uuid
}
