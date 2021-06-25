package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
悬挂链状态回传确认 APIRequest
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

func NewTaobaoWdkEquipmentConveyorStatusconfirmRequest() *TaobaoWdkEquipmentConveyorStatusconfirmRequest{
    return &TaobaoWdkEquipmentConveyorStatusconfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWdkEquipmentConveyorStatusconfirmRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.statusconfirm"
}

func (r TaobaoWdkEquipmentConveyorStatusconfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWdkEquipmentConveyorStatusconfirmRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r TaobaoWdkEquipmentConveyorStatusconfirmRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

func (r *TaobaoWdkEquipmentConveyorStatusconfirmRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r TaobaoWdkEquipmentConveyorStatusconfirmRequest) GetUuid() string {
    return r.uuid
}

