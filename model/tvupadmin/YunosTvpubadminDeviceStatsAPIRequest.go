package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindevicestatsAPIRequest 获取设备统计数据 API请求
// yunos.tvpubadmin.device.stats
//
// 获取设备统计数据
type YunostvpubadmindevicestatsAPIRequest struct {
	model.Params
	// 厂商名称
	_factoryName string
	// 设备型号
	_deviceModel string
}

// NewYunostvpubadmindevicestatsRequest 初始化YunostvpubadmindevicestatsAPIRequest对象
func NewYunostvpubadmindevicestatsRequest() *YunostvpubadmindevicestatsAPIRequest {
	return &YunostvpubadmindevicestatsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindevicestatsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.stats"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindevicestatsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindevicestatsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFactoryName is FactoryName Setter
// 厂商名称
func (r *YunostvpubadmindevicestatsAPIRequest) SetFactoryName(_factoryName string) error {
	r._factoryName = _factoryName
	r.Set("factory_name", _factoryName)
	return nil
}

// GetFactoryName FactoryName Getter
func (r YunostvpubadmindevicestatsAPIRequest) GetFactoryName() string {
	return r._factoryName
}

// SetDeviceModel is DeviceModel Setter
// 设备型号
func (r *YunostvpubadmindevicestatsAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// GetDeviceModel DeviceModel Getter
func (r YunostvpubadmindevicestatsAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}
