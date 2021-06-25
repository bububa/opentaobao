package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
容器被悬挂链扫描 APIRequest
taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor

容器被悬挂链扫描
*/
type TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest struct {
    model.Params

    // warehouse_code
    warehouseCode   string 

    // wcs_num
    wcsNum   int64 

}

func NewTaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest() *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest{
    return &TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor"
}

func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest) SetWcsNum(wcsNum int64) error {
    r.wcsNum = wcsNum
    r.Set("wcs_num", wcsNum)
    return nil
}

func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest) GetWcsNum() int64 {
    return r.wcsNum
}

