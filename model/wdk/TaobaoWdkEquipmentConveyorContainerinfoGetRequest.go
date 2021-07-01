package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取批次或波次中容器的信息 API请求
taobao.wdk.equipment.conveyor.containerinfo.get

获取批次或波次中容器的信息
*/
type TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest struct {
    model.Params
    // 仓库id
    _warehouseId   int64
    // 容器号
    _containerCode   string
    // 批次号，可以为空串
    _batchCode   string
    // 波次号，可以为空串
    _waveCode   string
}

// 初始化TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest对象
func NewTaobaoWdkEquipmentConveyorContainerinfoGetRequest() *TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest{
    return &TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.containerinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseId Setter
// 仓库id
func (r *TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) SetWarehouseId(_warehouseId int64) error {
    r._warehouseId = _warehouseId
    r.Set("warehouse_id", _warehouseId)
    return nil
}

// WarehouseId Getter
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetWarehouseId() int64 {
    return r._warehouseId
}
// ContainerCode Setter
// 容器号
func (r *TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) SetContainerCode(_containerCode string) error {
    r._containerCode = _containerCode
    r.Set("container_code", _containerCode)
    return nil
}

// ContainerCode Getter
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetContainerCode() string {
    return r._containerCode
}
// BatchCode Setter
// 批次号，可以为空串
func (r *TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) SetBatchCode(_batchCode string) error {
    r._batchCode = _batchCode
    r.Set("batch_code", _batchCode)
    return nil
}

// BatchCode Getter
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetBatchCode() string {
    return r._batchCode
}
// WaveCode Setter
// 波次号，可以为空串
func (r *TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) SetWaveCode(_waveCode string) error {
    r._waveCode = _waveCode
    r.Set("wave_code", _waveCode)
    return nil
}

// WaveCode Getter
func (r TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest) GetWaveCode() string {
    return r._waveCode
}
