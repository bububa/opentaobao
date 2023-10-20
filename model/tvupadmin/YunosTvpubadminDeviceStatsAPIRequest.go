package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceStatsAPIRequest 获取设备统计数据 API请求
// yunos.tvpubadmin.device.stats
//
// 获取设备统计数据
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceStatsAPIRequest) Reset() {
	r._factoryName = ""
	r._deviceModel = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceStatsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.stats"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceStatsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceStatsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFactoryName is FactoryName Setter
// 厂商名称
func (r *YunosTvpubadminDeviceStatsAPIRequest) SetFactoryName(_factoryName string) error {
	r._factoryName = _factoryName
	r.Set("factory_name", _factoryName)
	return nil
}

// GetFactoryName FactoryName Getter
func (r YunosTvpubadminDeviceStatsAPIRequest) GetFactoryName() string {
	return r._factoryName
}

// SetDeviceModel is DeviceModel Setter
// 设备型号
func (r *YunosTvpubadminDeviceStatsAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// GetDeviceModel DeviceModel Getter
func (r YunosTvpubadminDeviceStatsAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

var poolYunosTvpubadminDeviceStatsAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceStatsRequest()
	},
}

// GetYunosTvpubadminDeviceStatsRequest 从 sync.Pool 获取 YunosTvpubadminDeviceStatsAPIRequest
func GetYunosTvpubadminDeviceStatsAPIRequest() *YunosTvpubadminDeviceStatsAPIRequest {
	return poolYunosTvpubadminDeviceStatsAPIRequest.Get().(*YunosTvpubadminDeviceStatsAPIRequest)
}

// ReleaseYunosTvpubadminDeviceStatsAPIRequest 将 YunosTvpubadminDeviceStatsAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceStatsAPIRequest(v *YunosTvpubadminDeviceStatsAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceStatsAPIRequest.Put(v)
}
