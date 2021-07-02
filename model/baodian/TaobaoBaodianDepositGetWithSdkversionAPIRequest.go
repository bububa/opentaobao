package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaodianDepositGetWithSdkversionAPIRequest 查询用户宝点信息（带sdk版本，已迁移） API请求
// taobao.baodian.deposit.get.with.sdkversion
//
// 获取用户宝点信息（带sdk版本，已迁移）
type TaobaoBaodianDepositGetWithSdkversionAPIRequest struct {
	model.Params
	// 设备型号
	_deviceModel string
	// uuid
	_uuid string
	// sdk版本
	_sdkVersion string
}

// NewTaobaoBaodianDepositGetWithSdkversionRequest 初始化TaobaoBaodianDepositGetWithSdkversionAPIRequest对象
func NewTaobaoBaodianDepositGetWithSdkversionRequest() *TaobaoBaodianDepositGetWithSdkversionAPIRequest {
	return &TaobaoBaodianDepositGetWithSdkversionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaodianDepositGetWithSdkversionAPIRequest) GetApiMethodName() string {
	return "taobao.baodian.deposit.get.with.sdkversion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaodianDepositGetWithSdkversionAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDeviceModel is DeviceModel Setter
// 设备型号
func (r *TaobaoBaodianDepositGetWithSdkversionAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// GetDeviceModel DeviceModel Getter
func (r TaobaoBaodianDepositGetWithSdkversionAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

// SetUuid is Uuid Setter
// uuid
func (r *TaobaoBaodianDepositGetWithSdkversionAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaoBaodianDepositGetWithSdkversionAPIRequest) GetUuid() string {
	return r._uuid
}

// SetSdkVersion is SdkVersion Setter
// sdk版本
func (r *TaobaoBaodianDepositGetWithSdkversionAPIRequest) SetSdkVersion(_sdkVersion string) error {
	r._sdkVersion = _sdkVersion
	r.Set("sdk_version", _sdkVersion)
	return nil
}

// GetSdkVersion SdkVersion Getter
func (r TaobaoBaodianDepositGetWithSdkversionAPIRequest) GetSdkVersion() string {
	return r._sdkVersion
}
