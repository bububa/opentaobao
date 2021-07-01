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
type AlibabaRetailDeviceInventorySyncAPIRequest struct {
    model.Params
    // 设备类型
    _deviceType   string
    // 设备Id
    _deviceId   string
    // 系统自动生成
    _inventoryDtos   []InventorySyncDTO
    // 系统自动生成
    _deviceOption   *InventorySyncOption
}

// 初始化AlibabaRetailDeviceInventorySyncAPIRequest对象
func NewAlibabaRetailDeviceInventorySyncRequest() *AlibabaRetailDeviceInventorySyncAPIRequest{
    return &AlibabaRetailDeviceInventorySyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetApiMethodName() string {
    return "alibaba.retail.device.inventory.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceType Setter
// 设备类型
func (r *AlibabaRetailDeviceInventorySyncAPIRequest) SetDeviceType(_deviceType string) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetDeviceType() string {
    return r._deviceType
}
// DeviceId Setter
// 设备Id
func (r *AlibabaRetailDeviceInventorySyncAPIRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetDeviceId() string {
    return r._deviceId
}
// InventoryDtos Setter
// 系统自动生成
func (r *AlibabaRetailDeviceInventorySyncAPIRequest) SetInventoryDtos(_inventoryDtos []InventorySyncDTO) error {
    r._inventoryDtos = _inventoryDtos
    r.Set("inventory_dtos", _inventoryDtos)
    return nil
}

// InventoryDtos Getter
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetInventoryDtos() []InventorySyncDTO {
    return r._inventoryDtos
}
// DeviceOption Setter
// 系统自动生成
func (r *AlibabaRetailDeviceInventorySyncAPIRequest) SetDeviceOption(_deviceOption *InventorySyncOption) error {
    r._deviceOption = _deviceOption
    r.Set("device_option", _deviceOption)
    return nil
}

// DeviceOption Getter
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetDeviceOption() *InventorySyncOption {
    return r._deviceOption
}
