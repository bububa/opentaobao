package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretaildeviceinventorysyncAPIRequest 库存同步接口 API请求
// alibaba.retail.device.inventory.sync
//
// 商库存同步接口
type AlibabaretaildeviceinventorysyncAPIRequest struct {
	model.Params
	// 系统自动生成
	_inventoryDtos []InventorySyncDto
	// 设备类型
	_deviceType string
	// 设备Id
	_deviceId string
	// 系统自动生成
	_deviceOption *InventorySyncOption
}

// NewAlibabaretaildeviceinventorysyncRequest 初始化AlibabaretaildeviceinventorysyncAPIRequest对象
func NewAlibabaretaildeviceinventorysyncRequest() *AlibabaretaildeviceinventorysyncAPIRequest {
	return &AlibabaretaildeviceinventorysyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretaildeviceinventorysyncAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.device.inventory.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretaildeviceinventorysyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretaildeviceinventorysyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInventoryDtos is InventoryDtos Setter
// 系统自动生成
func (r *AlibabaretaildeviceinventorysyncAPIRequest) SetInventoryDtos(_inventoryDtos []InventorySyncDto) error {
	r._inventoryDtos = _inventoryDtos
	r.Set("inventory_dtos", _inventoryDtos)
	return nil
}

// GetInventoryDtos InventoryDtos Getter
func (r AlibabaretaildeviceinventorysyncAPIRequest) GetInventoryDtos() []InventorySyncDto {
	return r._inventoryDtos
}

// SetDeviceType is DeviceType Setter
// 设备类型
func (r *AlibabaretaildeviceinventorysyncAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r AlibabaretaildeviceinventorysyncAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetDeviceId is DeviceId Setter
// 设备Id
func (r *AlibabaretaildeviceinventorysyncAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaretaildeviceinventorysyncAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetDeviceOption is DeviceOption Setter
// 系统自动生成
func (r *AlibabaretaildeviceinventorysyncAPIRequest) SetDeviceOption(_deviceOption *InventorySyncOption) error {
	r._deviceOption = _deviceOption
	r.Set("device_option", _deviceOption)
	return nil
}

// GetDeviceOption DeviceOption Getter
func (r AlibabaretaildeviceinventorysyncAPIRequest) GetDeviceOption() *InventorySyncOption {
	return r._deviceOption
}
