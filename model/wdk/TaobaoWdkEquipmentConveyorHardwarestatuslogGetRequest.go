package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
硬件状态变化日志查询 API请求
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

// 初始化TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest对象
func NewTaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest() *TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest{
    return &TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.hardwarestatuslog.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseId Setter
// 仓库Id
func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) SetWarehouseId(warehouseId int64) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

// WarehouseId Getter
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) GetWarehouseId() int64 {
    return r.warehouseId
}
// ConveyorId Setter
// 悬挂链Id，即wcsNum
func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) SetConveyorId(conveyorId int64) error {
    r.conveyorId = conveyorId
    r.Set("conveyor_id", conveyorId)
    return nil
}

// ConveyorId Getter
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) GetConveyorId() int64 {
    return r.conveyorId
}
// StartId Setter
// 数据库id最小值
func (r *TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) SetStartId(startId int64) error {
    r.startId = startId
    r.Set("start_id", startId)
    return nil
}

// StartId Getter
func (r TaobaoWdkEquipmentConveyorHardwarestatuslogGetRequest) GetStartId() int64 {
    return r.startId
}
