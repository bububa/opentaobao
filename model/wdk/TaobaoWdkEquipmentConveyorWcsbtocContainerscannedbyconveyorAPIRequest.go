package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest 容器被悬挂链扫描 API请求
// taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor
//
// 容器被悬挂链扫描
type TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest struct {
	model.Params
	// warehouse_code
	_warehouseCode string
	// wcs_num
	_wcsNum int64
}

// NewTaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest 初始化TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest对象
func NewTaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest() *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest {
	return &TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// warehouse_code
func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetWcsNum is WcsNum Setter
// wcs_num
func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest) SetWcsNum(_wcsNum int64) error {
	r._wcsNum = _wcsNum
	r.Set("wcs_num", _wcsNum)
	return nil
}

// GetWcsNum WcsNum Getter
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest) GetWcsNum() int64 {
	return r._wcsNum
}
