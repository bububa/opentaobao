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
type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest struct {
    model.Params
    // warehouse_code
    _warehouseCode   string
    // wcs_num
    _wcsNum   int64
}

// 初始化TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest对象
func NewTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest() *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest{
    return &TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// warehouse_code
func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
// WcsNum Setter
// wcs_num
func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest) SetWcsNum(_wcsNum int64) error {
    r._wcsNum = _wcsNum
    r.Set("wcs_num", _wcsNum)
    return nil
}

// WcsNum Getter
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest) GetWcsNum() int64 {
    return r._wcsNum
}
