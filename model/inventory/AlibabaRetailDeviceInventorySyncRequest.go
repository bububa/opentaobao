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
    _deviceType   string
    // 设备Id
    _deviceId   string
    // 系统自动生成
    _inventoryDtos   []InventorySyncDTO
    // 系统自动生成
    _deviceOption   *InventorySyncOption
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
func (r *AlibabaRetailDeviceInventorySyncRequest) SetDeviceType(_deviceType string) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaRetailDeviceInventorySyncRequest) GetDeviceType() string {
    return r._deviceType
}
// DeviceId Setter
// 设备Id
func (r *AlibabaRetailDeviceInventorySyncRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaRetailDeviceInventorySyncRequest) GetDeviceId() string {
    return r._deviceId
}
// InventoryDtos Setter
// 系统自动生成
func (r *AlibabaRetailDeviceInventorySyncRequest) SetInventoryDtos(_inventoryDtos []InventorySyncDTO) error {
    r._inventoryDtos = _inventoryDtos
    r.Set("inventory_dtos", _inventoryDtos)
    return nil
}

// InventoryDtos Getter
func (r AlibabaRetailDeviceInventorySyncRequest) GetInventoryDtos() []InventorySyncDTO {
    return r._inventoryDtos
}
// DeviceOption Setter
// 系统自动生成
func (r *AlibabaRetailDeviceInventorySyncRequest) SetDeviceOption(_deviceOption *InventorySyncOption) error {
    r._deviceOption = _deviceOption
    r.Set("device_option", _deviceOption)
    return nil
}

// DeviceOption Getter
func (r AlibabaRetailDeviceInventorySyncRequest) GetDeviceOption() *InventorySyncOption {
    return r._deviceOption
}
