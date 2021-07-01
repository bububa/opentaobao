package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
容器被预分拣器分配到悬挂链 API请求
taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor

容器被预分拣器分配到悬挂链
*/
type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest struct {
    model.Params
    // warehouse_code
    _warehouseCode   string
    // wcs_num
    _wcsNum   int64
}

// 初始化TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest对象
func NewTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest() *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest{
    return &TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// warehouse_code
func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
// WcsNum Setter
// wcs_num
func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) SetWcsNum(_wcsNum int64) error {
    r._wcsNum = _wcsNum
    r.Set("wcs_num", _wcsNum)
    return nil
}

// WcsNum Getter
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest) GetWcsNum() int64 {
    return r._wcsNum
}
