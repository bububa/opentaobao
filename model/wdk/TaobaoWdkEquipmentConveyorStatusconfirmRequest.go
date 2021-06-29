package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
悬挂链状态回传确认 API请求
taobao.wdk.equipment.conveyor.statusconfirm

悬挂链状态回传确认
*/
type TaobaoWdkEquipmentConveyorStatusconfirmRequest struct {
    model.Params
    // warehouse_code
    warehouseCode   string
    // uuid
    uuid   string
}

// 初始化TaobaoWdkEquipmentConveyorStatusconfirmRequest对象
func NewTaobaoWdkEquipmentConveyorStatusconfirmRequest() *TaobaoWdkEquipmentConveyorStatusconfirmRequest{
    return &TaobaoWdkEquipmentConveyorStatusconfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorStatusconfirmRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.statusconfirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorStatusconfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// warehouse_code
func (r *TaobaoWdkEquipmentConveyorStatusconfirmRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorStatusconfirmRequest) GetWarehouseCode() string {
    return r.warehouseCode
}
// Uuid Setter
// uuid
func (r *TaobaoWdkEquipmentConveyorStatusconfirmRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r TaobaoWdkEquipmentConveyorStatusconfirmRequest) GetUuid() string {
    return r.uuid
}
