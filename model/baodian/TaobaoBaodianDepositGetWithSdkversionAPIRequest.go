package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaodiandepositgetwithsdkversionAPIRequest 查询用户宝点信息（带sdk版本，已迁移） API请求
// taobao.baodian.deposit.get.with.sdkversion
//
// 获取用户宝点信息（带sdk版本，已迁移）
type TaobaobaodiandepositgetwithsdkversionAPIRequest struct {
	model.Params
	// 设备型号
	_deviceModel string
	// uuid
	_uuid string
	// sdk版本
	_sdkVersion string
}

// NewTaobaobaodiandepositgetwithsdkversionRequest 初始化TaobaobaodiandepositgetwithsdkversionAPIRequest对象
func NewTaobaobaodiandepositgetwithsdkversionRequest() *TaobaobaodiandepositgetwithsdkversionAPIRequest {
	return &TaobaobaodiandepositgetwithsdkversionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaodiandepositgetwithsdkversionAPIRequest) GetApiMethodName() string {
	return "taobao.baodian.deposit.get.with.sdkversion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaodiandepositgetwithsdkversionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaodiandepositgetwithsdkversionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceModel is DeviceModel Setter
// 设备型号
func (r *TaobaobaodiandepositgetwithsdkversionAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// GetDeviceModel DeviceModel Getter
func (r TaobaobaodiandepositgetwithsdkversionAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

// SetUuid is Uuid Setter
// uuid
func (r *TaobaobaodiandepositgetwithsdkversionAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaobaodiandepositgetwithsdkversionAPIRequest) GetUuid() string {
	return r._uuid
}

// SetSdkVersion is SdkVersion Setter
// sdk版本
func (r *TaobaobaodiandepositgetwithsdkversionAPIRequest) SetSdkVersion(_sdkVersion string) error {
	r._sdkVersion = _sdkVersion
	r.Set("sdk_version", _sdkVersion)
	return nil
}

// GetSdkVersion SdkVersion Getter
func (r TaobaobaodiandepositgetwithsdkversionAPIRequest) GetSdkVersion() string {
	return r._sdkVersion
}
