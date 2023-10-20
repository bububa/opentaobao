package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest 容器被预分拣器分配到悬挂链 API请求
// taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor
//
// 容器被预分拣器分配到悬挂链
type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest struct {
	model.Params
	// warehouse_code
	_warehouseCode string
	// wcs_num
	_wcsNum int64
}

// NewTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest 初始化TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest对象
func NewTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest() *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest {
	return &TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) Reset() {
	r._warehouseCode = ""
	r._wcsNum = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// warehouse_code
func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetWcsNum is WcsNum Setter
// wcs_num
func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) SetWcsNum(_wcsNum int64) error {
	r._wcsNum = _wcsNum
	r.Set("wcs_num", _wcsNum)
	return nil
}

// GetWcsNum WcsNum Getter
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) GetWcsNum() int64 {
	return r._wcsNum
}

var poolTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest()
	},
}

// GetTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest 从 sync.Pool 获取 TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest
func GetTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest() *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest {
	return poolTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest.Get().(*TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest)
}

// ReleaseTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest 将 TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest 放入 sync.Pool
func ReleaseTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest(v *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) {
	v.Reset()
	poolTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest.Put(v)
}
