package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取悬挂链系统事件 API请求
taobao.wdk.equipment.conveyor.systemevent.get

五道口悬挂链系统事件查询
*/
type TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest struct {
    model.Params
    // 仓库Id
    _warehouseId   int64
    // 悬挂链Id，即wcsNum
    _conveyorId   int64
    // 数据库id最小值
    _startId   int64
}

// 初始化TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest对象
func NewTaobaoWdkEquipmentConveyorSystemeventGetRequest() *TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest{
    return &TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.systemevent.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseId Setter
// 仓库Id
func (r *TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) SetWarehouseId(_warehouseId int64) error {
    r._warehouseId = _warehouseId
    r.Set("warehouse_id", _warehouseId)
    return nil
}

// WarehouseId Getter
func (r TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) GetWarehouseId() int64 {
    return r._warehouseId
}
// ConveyorId Setter
// 悬挂链Id，即wcsNum
func (r *TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) SetConveyorId(_conveyorId int64) error {
    r._conveyorId = _conveyorId
    r.Set("conveyor_id", _conveyorId)
    return nil
}

// ConveyorId Getter
func (r TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) GetConveyorId() int64 {
    return r._conveyorId
}
// StartId Setter
// 数据库id最小值
func (r *TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) SetStartId(_startId int64) error {
    r._startId = _startId
    r.Set("start_id", _startId)
    return nil
}

// StartId Getter
func (r TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest) GetStartId() int64 {
    return r._startId
}
