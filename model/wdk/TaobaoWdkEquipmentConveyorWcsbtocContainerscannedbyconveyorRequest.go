package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
容器被悬挂链扫描 API请求
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

// 初始化TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest对象
func NewTaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest() *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest{
    return &TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// warehouse_code
func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest) GetWarehouseCode() string {
    return r.warehouseCode
}
// WcsNum Setter
// wcs_num
func (r *TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest) SetWcsNum(wcsNum int64) error {
    r.wcsNum = wcsNum
    r.Set("wcs_num", wcsNum)
    return nil
}

// WcsNum Getter
func (r TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorRequest) GetWcsNum() int64 {
    return r.wcsNum
}
