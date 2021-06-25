package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
硬件状态变化日志查询 APIRequest
taobao.wdk.equipment.conveyor.hardwarestatuslog.get

硬件状态变化日志查询
*/
type TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest struct {
    model.Params

    // 仓库Id
    warehouseId   int64 

    // 悬挂链Id，即wcsNum
    conveyorId   int64 

    // 数据库id最小值
    startId   int64 

}

func NewTaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest() *TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest{
    return &TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.hardwarestatuslog.get"
}

func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) SetWarehouseId(warehouseId int64) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) GetWarehouseId() int64 {
    return r.warehouseId
}

func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) SetConveyorId(conveyorId int64) error {
    r.conveyorId = conveyorId
    r.Set("conveyor_id", conveyorId)
    return nil
}

func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) GetConveyorId() int64 {
    return r.conveyorId
}

func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) SetStartId(startId int64) error {
    r.startId = startId
    r.Set("start_id", startId)
    return nil
}

func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) GetStartId() int64 {
    return r.startId
}

