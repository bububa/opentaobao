package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
异常通道日志查询 API请求
taobao.wdk.equipment.conveyor.exceptionslidewaylog.get

五道口悬挂链异常通道事件查询
*/
type TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest struct {
    model.Params
    // 仓库Id
    warehouseId   int64
    // 悬挂链Id，即wcsNum
    conveyorId   int64
    // 数据库id最小值
    startId   int64
}

// 初始化TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest对象
func NewTaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest() *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest{
    return &TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.exceptionslidewaylog.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseId Setter
// 仓库Id
func (r *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest) SetWarehouseId(warehouseId int64) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

// WarehouseId Getter
func (r TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest) GetWarehouseId() int64 {
    return r.warehouseId
}
// ConveyorId Setter
// 悬挂链Id，即wcsNum
func (r *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest) SetConveyorId(conveyorId int64) error {
    r.conveyorId = conveyorId
    r.Set("conveyor_id", conveyorId)
    return nil
}

// ConveyorId Getter
func (r TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest) GetConveyorId() int64 {
    return r.conveyorId
}
// StartId Setter
// 数据库id最小值
func (r *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest) SetStartId(startId int64) error {
    r.startId = startId
    r.Set("start_id", startId)
    return nil
}

// StartId Getter
func (r TaobaoWdkEquipmentConveyorExceptionslidewaylogGetRequest) GetStartId() int64 {
    return r.startId
}
