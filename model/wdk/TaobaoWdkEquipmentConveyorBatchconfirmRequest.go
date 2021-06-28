package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口悬挂链信息批量确认 APIRequest
taobao.wdk.equipment.conveyor.batchconfirm

批量消息确认
*/
type TaobaoWdkEquipmentConveyorBatchconfirmRequest struct {
    model.Params

    // 仓库code
    warehouseCode   string 

    // 待确认的uuid列表
    uuids   []string 

}

func NewTaobaoWdkEquipmentConveyorBatchconfirmRequest() *TaobaoWdkEquipmentConveyorBatchconfirmRequest{
    return &TaobaoWdkEquipmentConveyorBatchconfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWdkEquipmentConveyorBatchconfirmRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.batchconfirm"
}

func (r TaobaoWdkEquipmentConveyorBatchconfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWdkEquipmentConveyorBatchconfirmRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r TaobaoWdkEquipmentConveyorBatchconfirmRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

func (r *TaobaoWdkEquipmentConveyorBatchconfirmRequest) SetUuids(uuids []string) error {
    r.uuids = uuids
    r.Set("uuids", uuids)
    return nil
}

func (r TaobaoWdkEquipmentConveyorBatchconfirmRequest) GetUuids() []string {
    return r.uuids
}

