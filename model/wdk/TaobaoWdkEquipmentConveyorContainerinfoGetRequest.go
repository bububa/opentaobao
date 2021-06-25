package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取批次或波次中容器的信息 APIRequest
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

func NewTaobaoWdkEquipmentConveyorContainerinfoGetRequest() *TaobaoWdkEquipmentConveyorContainerinfoGetRequest{
    return &TaobaoWdkEquipmentConveyorContainerinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWdkEquipmentConveyorContainerinfoGetRequest) GetApiMethodName() string {
    return "taobao.wdk.equipment.conveyor.containerinfo.get"
}

func (r TaobaoWdkEquipmentConveyorContainerinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWdkEquipmentConveyorContainerinfoGetRequest) SetWarehouseId(warehouseId int64) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

func (r TaobaoWdkEquipmentConveyorContainerinfoGetRequest) GetWarehouseId() int64 {
    return r.warehouseId
}

func (r *TaobaoWdkEquipmentConveyorContainerinfoGetRequest) SetContainerCode(containerCode string) error {
    r.containerCode = containerCode
    r.Set("container_code", containerCode)
    return nil
}

func (r TaobaoWdkEquipmentConveyorContainerinfoGetRequest) GetContainerCode() string {
    return r.containerCode
}

func (r *TaobaoWdkEquipmentConveyorContainerinfoGetRequest) SetBatchCode(batchCode string) error {
    r.batchCode = batchCode
    r.Set("batch_code", batchCode)
    return nil
}

func (r TaobaoWdkEquipmentConveyorContainerinfoGetRequest) GetBatchCode() string {
    return r.batchCode
}

func (r *TaobaoWdkEquipmentConveyorContainerinfoGetRequest) SetWaveCode(waveCode string) error {
    r.waveCode = waveCode
    r.Set("wave_code", waveCode)
    return nil
}

func (r TaobaoWdkEquipmentConveyorContainerinfoGetRequest) GetWaveCode() string {
    return r.waveCode
}

