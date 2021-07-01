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
type TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest struct {
    model.Params
    // warehouse_code
    _warehouseCode   string
    // uuid
    _uuid   string
}

// 初始化TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest对象
func NewTaobaoWdkEquipmentConveyorStatusconfirmRequest() *TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest{
    return &TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.statusconfirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// warehouse_code
func (r *TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
// Uuid Setter
// uuid
func (r *TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest) GetUuid() string {
    return r._uuid
}
