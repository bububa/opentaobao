package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest 获取五道口设备管理信息 API请求
// taobao.wdk.equipment.deviceadmin.deviceinfo.get
//
// 通过仓编码获取五道口设备管理信息
type TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest struct {
	model.Params
	// 仓编码
	_warehouseCode string
	// 设备类型
	_deviceType int64
}

// NewTaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest 初始化TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest对象
func NewTaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest() *TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest {
	return &TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest) Reset() {
	r._warehouseCode = ""
	r._deviceType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.equipment.deviceadmin.deviceinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 仓编码
func (r *TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetDeviceType is DeviceType Setter
// 设备类型
func (r *TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest) SetDeviceType(_deviceType int64) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest) GetDeviceType() int64 {
	return r._deviceType
}

var poolTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest()
	},
}

// GetTaobaoWdkEquipmentDeviceadminDeviceinfoGetRequest 从 sync.Pool 获取 TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest
func GetTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest() *TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest {
	return poolTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest.Get().(*TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest)
}

// ReleaseTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest 将 TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest(v *TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest) {
	v.Reset()
	poolTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest.Put(v)
}
