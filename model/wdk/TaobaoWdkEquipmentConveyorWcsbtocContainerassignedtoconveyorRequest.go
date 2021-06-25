package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
容器被预分拣器分配到悬挂链 APIRequest
taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor

容器被预分拣器分配到悬挂链
*/
type TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest struct {
    model.Params

    // warehouse_code
    warehouseCode   string 

    // wcs_num
    wcsNum   int64 

}

func NewTaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest() *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest{
    return &TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor"
}

func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest) SetWcsNum(wcsNum int64) error {
    r.wcsNum = wcsNum
    r.Set("wcs_num", wcsNum)
    return nil
}

func (r TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorRequest) GetWcsNum() int64 {
    return r.wcsNum
}

