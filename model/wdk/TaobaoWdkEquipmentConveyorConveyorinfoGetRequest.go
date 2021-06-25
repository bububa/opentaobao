package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取五道口悬挂链信息 APIRequest
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

func NewTaobaoWdkEquipmentConveyorConveyorinfoGetRequest() *TaobaoWdkEquipmentConveyorConveyorinfoGetRequest{
    return &TaobaoWdkEquipmentConveyorConveyorinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWdkEquipmentConveyorConveyorinfoGetRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.conveyorinfo.get"
}

func (r TaobaoWdkEquipmentConveyorConveyorinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWdkEquipmentConveyorConveyorinfoGetRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r TaobaoWdkEquipmentConveyorConveyorinfoGetRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

func (r *TaobaoWdkEquipmentConveyorConveyorinfoGetRequest) SetConveyorId(conveyorId int64) error {
    r.conveyorId = conveyorId
    r.Set("conveyor_id", conveyorId)
    return nil
}

func (r TaobaoWdkEquipmentConveyorConveyorinfoGetRequest) GetConveyorId() int64 {
    return r.conveyorId
}

