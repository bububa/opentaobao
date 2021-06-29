package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口悬挂链信息批量确认 API请求
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

// 初始化TaobaoWdkEquipmentConveyorBatchconfirmRequest对象
func NewTaobaoWdkEquipmentConveyorBatchconfirmRequest() *TaobaoWdkEquipmentConveyorBatchconfirmRequest{
    return &TaobaoWdkEquipmentConveyorBatchconfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorBatchconfirmRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.batchconfirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorBatchconfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 仓库code
func (r *TaobaoWdkEquipmentConveyorBatchconfirmRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorBatchconfirmRequest) GetWarehouseCode() string {
    return r.warehouseCode
}
// Uuids Setter
// 待确认的uuid列表
func (r *TaobaoWdkEquipmentConveyorBatchconfirmRequest) SetUuids(uuids []string) error {
    r.uuids = uuids
    r.Set("uuids", uuids)
    return nil
}

// Uuids Getter
func (r TaobaoWdkEquipmentConveyorBatchconfirmRequest) GetUuids() []string {
    return r.uuids
}
