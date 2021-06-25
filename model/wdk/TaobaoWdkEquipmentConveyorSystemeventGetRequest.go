package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取悬挂链系统事件 APIRequest
taobao.wdk.equipment.conveyor.systemevent.get

五道口悬挂链系统事件查询
*/
type TaobaoWdkEquipmentConveyorSystemeventGetRequest struct {
    model.Params

    // 仓库Id
    warehouseId   int64 

    // 悬挂链Id，即wcsNum
    conveyorId   int64 

    // 数据库id最小值
    startId   int64 

}

func NewTaobaoWdkEquipmentConveyorSystemeventGetRequest() *TaobaoWdkEquipmentConveyorSystemeventGetRequest{
    return &TaobaoWdkEquipmentConveyorSystemeventGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWdkEquipmentConveyorSystemeventGetRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.systemevent.get"
}

func (r TaobaoWdkEquipmentConveyorSystemeventGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWdkEquipmentConveyorSystemeventGetRequest) SetWarehouseId(warehouseId int64) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

func (r TaobaoWdkEquipmentConveyorSystemeventGetRequest) GetWarehouseId() int64 {
    return r.warehouseId
}

func (r *TaobaoWdkEquipmentConveyorSystemeventGetRequest) SetConveyorId(conveyorId int64) error {
    r.conveyorId = conveyorId
    r.Set("conveyor_id", conveyorId)
    return nil
}

func (r TaobaoWdkEquipmentConveyorSystemeventGetRequest) GetConveyorId() int64 {
    return r.conveyorId
}

func (r *TaobaoWdkEquipmentConveyorSystemeventGetRequest) SetStartId(startId int64) error {
    r.startId = startId
    r.Set("start_id", startId)
    return nil
}

func (r TaobaoWdkEquipmentConveyorSystemeventGetRequest) GetStartId() int64 {
    return r.startId
}

