package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest 获取五道口设备管理信息 API请求
// taobao.wdk.equipment.deviceadmin.deviceinfo.get
//
// 通过仓编码获取五道口设备管理信息
type TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest struct {
	model.Params
	// 仓编码
	_warehouseCode string
	// 设备类型
	_deviceType int64
}

// NewTaobaowdkequipmentdeviceadmindeviceinfogetRequest 初始化TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest对象
func NewTaobaowdkequipmentdeviceadmindeviceinfogetRequest() *TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest {
	return &TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.deviceadmin.deviceinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 仓编码
func (r *TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetDeviceType is DeviceType Setter
// 设备类型
func (r *TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest) SetDeviceType(_deviceType int64) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaowdkequipmentdeviceadmindeviceinfogetAPIRequest) GetDeviceType() int64 {
	return r._deviceType
}
