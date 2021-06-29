package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存同步接口 API请求
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

// 初始化AlibabaRetailDeviceInventorySyncRequest对象
func NewAlibabaRetailDeviceInventorySyncRequest() *AlibabaRetailDeviceInventorySyncRequest{
    return &AlibabaRetailDeviceInventorySyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailDeviceInventorySyncRequest) GetApiMethodName() string {
    return "alibaba.retail.device.inventory.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailDeviceInventorySyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceType Setter
// 设备类型
func (r *AlibabaRetailDeviceInventorySyncRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaRetailDeviceInventorySyncRequest) GetDeviceType() string {
    return r.deviceType
}
// DeviceId Setter
// 设备Id
func (r *AlibabaRetailDeviceInventorySyncRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaRetailDeviceInventorySyncRequest) GetDeviceId() string {
    return r.deviceId
}
// InventoryDtos Setter
// 系统自动生成
func (r *AlibabaRetailDeviceInventorySyncRequest) SetInventoryDtos(inventoryDtos []InventorySyncDTO) error {
    r.inventoryDtos = inventoryDtos
    r.Set("inventory_dtos", inventoryDtos)
    return nil
}

// InventoryDtos Getter
func (r AlibabaRetailDeviceInventorySyncRequest) GetInventoryDtos() []InventorySyncDTO {
    return r.inventoryDtos
}
// DeviceOption Setter
// 系统自动生成
func (r *AlibabaRetailDeviceInventorySyncRequest) SetDeviceOption(deviceOption *InventorySyncOption) error {
    r.deviceOption = deviceOption
    r.Set("device_option", deviceOption)
    return nil
}

// DeviceOption Getter
func (r AlibabaRetailDeviceInventorySyncRequest) GetDeviceOption() *InventorySyncOption {
    return r.deviceOption
}
