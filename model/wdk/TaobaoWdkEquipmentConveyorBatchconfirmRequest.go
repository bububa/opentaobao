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
type TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest struct {
    model.Params
    // 仓库code
    _warehouseCode   string
    // 待确认的uuid列表
    _uuids   []string
}

// 初始化TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest对象
func NewTaobaoWdkEquipmentConveyorBatchconfirmRequest() *TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest{
    return &TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.batchconfirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 仓库code
func (r *TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
// Uuids Setter
// 待确认的uuid列表
func (r *TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) SetUuids(_uuids []string) error {
    r._uuids = _uuids
    r.Set("uuids", _uuids)
    return nil
}

// Uuids Getter
func (r TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest) GetUuids() []string {
    return r._uuids
}
