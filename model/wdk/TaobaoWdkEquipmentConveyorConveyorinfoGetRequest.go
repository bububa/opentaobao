package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取五道口悬挂链信息 API请求
taobao.wdk.equipment.conveyor.conveyorinfo.get

获取五道口悬挂链信息
*/
type TaobaoWdkEquipmentConveyorConveyorinfoGetRequest struct {
    model.Params
    // 仓库code
    warehouseCode   string
    // wcsNum
    conveyorId   int64
}

// 初始化TaobaoWdkEquipmentConveyorConveyorinfoGetRequest对象
func NewTaobaoWdkEquipmentConveyorConveyorinfoGetRequest() *TaobaoWdkEquipmentConveyorConveyorinfoGetRequest{
    return &TaobaoWdkEquipmentConveyorConveyorinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.conveyorinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 仓库code
func (r *TaobaoWdkEquipmentConveyorConveyorinfoGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}
// ConveyorId Setter
// wcsNum
func (r *TaobaoWdkEquipmentConveyorConveyorinfoGetRequest) SetConveyorId(conveyorId int64) error {
    r.conveyorId = conveyorId
    r.Set("conveyor_id", conveyorId)
    return nil
}

// ConveyorId Getter
func (r TaobaoWdkEquipmentConveyorConveyorinfoGetRequest) GetConveyorId() int64 {
    return r.conveyorId
}
