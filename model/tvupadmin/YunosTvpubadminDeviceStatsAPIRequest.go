package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceStatsAPIRequest
获取设备统计数据 API请求
yunos.tvpubadmin.device.stats

获取设备统计数据 */
type YunosTvpubadminDeviceStatsAPIRequest struct {
	model.Params
	// 厂商名称
	_factoryName string
	// 设备型号
	_deviceModel string
}

// NewYunosTvpubadminDeviceStatsRequest 初始化YunosTvpubadminDeviceStatsAPIRequest对象
func NewYunosTvpubadminDeviceStatsRequest() *YunosTvpubadminDeviceStatsAPIRequest {
	return &YunosTvpubadminDeviceStatsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceStatsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.stats"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceStatsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is FactoryName Setter
// 厂商名称
func (r *YunosTvpubadminDeviceStatsAPIRequest) SetFactoryName(_factoryName string) error {
	r._factoryName = _factoryName
	r.Set("factory_name", _factoryName)
	return nil
}

// Get FactoryName Getter
func (r YunosTvpubadminDeviceStatsAPIRequest) GetFactoryName() string {
	return r._factoryName
}

// Set is DeviceModel Setter
// 设备型号
func (r *YunosTvpubadminDeviceStatsAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// Get DeviceModel Getter
func (r YunosTvpubadminDeviceStatsAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}
