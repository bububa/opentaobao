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
    _warehouseCode   string
    // uuid
    _uuid   string
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
func (r *TaobaoWdkEquipmentConveyorStatusconfirmRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorStatusconfirmRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
// Uuid Setter
// uuid
func (r *TaobaoWdkEquipmentConveyorStatusconfirmRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r TaobaoWdkEquipmentConveyorStatusconfirmRequest) GetUuid() string {
    return r._uuid
}
