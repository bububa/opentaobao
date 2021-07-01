package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaodianDepositGetWithSdkversionAPIRequest
查询用户宝点信息（带sdk版本，已迁移） API请求
taobao.baodian.deposit.get.with.sdkversion

获取用户宝点信息（带sdk版本，已迁移） */
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

// Set is DeviceModel Setter
// 设备型号
func (r *TaobaoBaodianDepositGetWithSdkversionAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// Get DeviceModel Getter
func (r TaobaoBaodianDepositGetWithSdkversionAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

// Set is Uuid Setter
// uuid
func (r *TaobaoBaodianDepositGetWithSdkversionAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r TaobaoBaodianDepositGetWithSdkversionAPIRequest) GetUuid() string {
	return r._uuid
}

// Set is SdkVersion Setter
// sdk版本
func (r *TaobaoBaodianDepositGetWithSdkversionAPIRequest) SetSdkVersion(_sdkVersion string) error {
	r._sdkVersion = _sdkVersion
	r.Set("sdk_version", _sdkVersion)
	return nil
}

// Get SdkVersion Getter
func (r TaobaoBaodianDepositGetWithSdkversionAPIRequest) GetSdkVersion() string {
	return r._sdkVersion
}
