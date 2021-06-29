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
type TaobaoWdkEquipmentConveyorContainerinfoGetRequest struct {
    model.Params
    // 仓库id
    warehouseId   int64
    // 容器号
    containerCode   string
    // 批次号，可以为空串
    batchCode   string
    // 波次号，可以为空串
    waveCode   string
}

// 初始化TaobaoWdkEquipmentConveyorContainerinfoGetRequest对象
func NewTaobaoWdkEquipmentConveyorContainerinfoGetRequest() *TaobaoWdkEquipmentConveyorContainerinfoGetRequest{
    return &TaobaoWdkEquipmentConveyorContainerinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentConveyorContainerinfoGetRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.containerinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentConveyorContainerinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseId Setter
// 仓库id
func (r *TaobaoWdkEquipmentConveyorContainerinfoGetRequest) SetWarehouseId(warehouseId int64) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

// WarehouseId Getter
func (r TaobaoWdkEquipmentConveyorContainerinfoGetRequest) GetWarehouseId() int64 {
    return r.warehouseId
}
// ContainerCode Setter
// 容器号
func (r *TaobaoWdkEquipmentConveyorContainerinfoGetRequest) SetContainerCode(containerCode string) error {
    r.containerCode = containerCode
    r.Set("container_code", containerCode)
    return nil
}

// ContainerCode Getter
func (r TaobaoWdkEquipmentConveyorContainerinfoGetRequest) GetContainerCode() string {
    return r.containerCode
}
// BatchCode Setter
// 批次号，可以为空串
func (r *TaobaoWdkEquipmentConveyorContainerinfoGetRequest) SetBatchCode(batchCode string) error {
    r.batchCode = batchCode
    r.Set("batch_code", batchCode)
    return nil
}

// BatchCode Getter
func (r TaobaoWdkEquipmentConveyorContainerinfoGetRequest) GetBatchCode() string {
    return r.batchCode
}
// WaveCode Setter
// 波次号，可以为空串
func (r *TaobaoWdkEquipmentConveyorContainerinfoGetRequest) SetWaveCode(waveCode string) error {
    r.waveCode = waveCode
    r.Set("wave_code", waveCode)
    return nil
}

// WaveCode Getter
func (r TaobaoWdkEquipmentConveyorContainerinfoGetRequest) GetWaveCode() string {
    return r.waveCode
}
