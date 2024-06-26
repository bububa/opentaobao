package inventory

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailDeviceInventorySyncAPIRequest 库存同步接口 API请求
// alibaba.retail.device.inventory.sync
//
// 商库存同步接口
type AlibabaRetailDeviceInventorySyncAPIRequest struct {
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

// NewAlibabaRetailDeviceInventorySyncRequest 初始化AlibabaRetailDeviceInventorySyncAPIRequest对象
func NewAlibabaRetailDeviceInventorySyncRequest() *AlibabaRetailDeviceInventorySyncAPIRequest {
	return &AlibabaRetailDeviceInventorySyncAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailDeviceInventorySyncAPIRequest) Reset() {
	r._inventoryDtos = r._inventoryDtos[:0]
	r._deviceType = ""
	r._deviceId = ""
	r._deviceOption = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.device.inventory.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInventoryDtos is InventoryDtos Setter
// 系统自动生成
func (r *AlibabaRetailDeviceInventorySyncAPIRequest) SetInventoryDtos(_inventoryDtos []InventorySyncDto) error {
	r._inventoryDtos = _inventoryDtos
	r.Set("inventory_dtos", _inventoryDtos)
	return nil
}

// GetInventoryDtos InventoryDtos Getter
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetInventoryDtos() []InventorySyncDto {
	return r._inventoryDtos
}

// SetDeviceType is DeviceType Setter
// 设备类型
func (r *AlibabaRetailDeviceInventorySyncAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetDeviceId is DeviceId Setter
// 设备Id
func (r *AlibabaRetailDeviceInventorySyncAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetDeviceOption is DeviceOption Setter
// 系统自动生成
func (r *AlibabaRetailDeviceInventorySyncAPIRequest) SetDeviceOption(_deviceOption *InventorySyncOption) error {
	r._deviceOption = _deviceOption
	r.Set("device_option", _deviceOption)
	return nil
}

// GetDeviceOption DeviceOption Getter
func (r AlibabaRetailDeviceInventorySyncAPIRequest) GetDeviceOption() *InventorySyncOption {
	return r._deviceOption
}

var poolAlibabaRetailDeviceInventorySyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailDeviceInventorySyncRequest()
	},
}

// GetAlibabaRetailDeviceInventorySyncRequest 从 sync.Pool 获取 AlibabaRetailDeviceInventorySyncAPIRequest
func GetAlibabaRetailDeviceInventorySyncAPIRequest() *AlibabaRetailDeviceInventorySyncAPIRequest {
	return poolAlibabaRetailDeviceInventorySyncAPIRequest.Get().(*AlibabaRetailDeviceInventorySyncAPIRequest)
}

// ReleaseAlibabaRetailDeviceInventorySyncAPIRequest 将 AlibabaRetailDeviceInventorySyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailDeviceInventorySyncAPIRequest(v *AlibabaRetailDeviceInventorySyncAPIRequest) {
	v.Reset()
	poolAlibabaRetailDeviceInventorySyncAPIRequest.Put(v)
}
