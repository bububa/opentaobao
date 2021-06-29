package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存同步接口 APIRequest
alibaba.retail.device.inventory.sync

商库存同步接口
*/
type AlibabaRetailDeviceInventorySyncRequest struct {
    model.Params

    // 设备类型
    deviceType   string 

    // 设备Id
    deviceId   string 

    // 系统自动生成
    inventoryDtos   []InventorySyncDTO 

    // 系统自动生成
    deviceOption   *InventorySyncOption 

}

func NewAlibabaRetailDeviceInventorySyncRequest() *AlibabaRetailDeviceInventorySyncRequest{
    return &AlibabaRetailDeviceInventorySyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailDeviceInventorySyncRequest) GetApiMethodName() string {
    return "alibaba.retail.device.inventory.sync"
}

func (r AlibabaRetailDeviceInventorySyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailDeviceInventorySyncRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

func (r AlibabaRetailDeviceInventorySyncRequest) GetDeviceType() string {
    return r.deviceType
}

func (r *AlibabaRetailDeviceInventorySyncRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r AlibabaRetailDeviceInventorySyncRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *AlibabaRetailDeviceInventorySyncRequest) SetInventoryDtos(inventoryDtos []InventorySyncDTO) error {
    r.inventoryDtos = inventoryDtos
    r.Set("inventory_dtos", inventoryDtos)
    return nil
}

func (r AlibabaRetailDeviceInventorySyncRequest) GetInventoryDtos() []InventorySyncDTO {
    return r.inventoryDtos
}

func (r *AlibabaRetailDeviceInventorySyncRequest) SetDeviceOption(deviceOption *InventorySyncOption) error {
    r.deviceOption = deviceOption
    r.Set("device_option", deviceOption)
    return nil
}

func (r AlibabaRetailDeviceInventorySyncRequest) GetDeviceOption() *InventorySyncOption {
    return r.deviceOption
}

