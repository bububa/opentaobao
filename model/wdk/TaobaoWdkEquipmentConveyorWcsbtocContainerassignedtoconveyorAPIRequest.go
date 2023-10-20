package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest 容器被预分拣器分配到悬挂链 API请求
// taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor
//
// 容器被预分拣器分配到悬挂链
type TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest struct {
	model.Params
	// warehouse_code
	_warehouseCode string
	// wcs_num
	_wcsNum int64
}

// NewTaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorRequest 初始化TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest对象
func NewTaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorRequest() *TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest {
	return &TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// warehouse_code
func (r *TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetWcsNum is WcsNum Setter
// wcs_num
func (r *TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest) SetWcsNum(_wcsNum int64) error {
	r._wcsNum = _wcsNum
	r.Set("wcs_num", _wcsNum)
	return nil
}

// GetWcsNum WcsNum Getter
func (r TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest) GetWcsNum() int64 {
	return r._wcsNum
}
