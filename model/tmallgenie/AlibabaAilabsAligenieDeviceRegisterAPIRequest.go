package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieDeviceRegisterAPIRequest 天猫精灵开放平台获取设备秘钥 API请求
// alibaba.ailabs.aligenie.device.register
//
// 向天猫精灵inside平台注册设备mac地址，并获取设备的唯一密钥
type AlibabaAilabsAligenieDeviceRegisterAPIRequest struct {
	model.Params
	// mac区段脚本
	_macSections string
	// 设备id
	_deviceId int64
}

// NewAlibabaAilabsAligenieDeviceRegisterRequest 初始化AlibabaAilabsAligenieDeviceRegisterAPIRequest对象
func NewAlibabaAilabsAligenieDeviceRegisterRequest() *AlibabaAilabsAligenieDeviceRegisterAPIRequest {
	return &AlibabaAilabsAligenieDeviceRegisterAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsAligenieDeviceRegisterAPIRequest) Reset() {
	r._macSections = ""
	r._deviceId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieDeviceRegisterAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.device.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieDeviceRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsAligenieDeviceRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMacSections is MacSections Setter
// mac区段脚本
func (r *AlibabaAilabsAligenieDeviceRegisterAPIRequest) SetMacSections(_macSections string) error {
	r._macSections = _macSections
	r.Set("mac_sections", _macSections)
	return nil
}

// GetMacSections MacSections Getter
func (r AlibabaAilabsAligenieDeviceRegisterAPIRequest) GetMacSections() string {
	return r._macSections
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *AlibabaAilabsAligenieDeviceRegisterAPIRequest) SetDeviceId(_deviceId int64) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaAilabsAligenieDeviceRegisterAPIRequest) GetDeviceId() int64 {
	return r._deviceId
}

var poolAlibabaAilabsAligenieDeviceRegisterAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsAligenieDeviceRegisterRequest()
	},
}

// GetAlibabaAilabsAligenieDeviceRegisterRequest 从 sync.Pool 获取 AlibabaAilabsAligenieDeviceRegisterAPIRequest
func GetAlibabaAilabsAligenieDeviceRegisterAPIRequest() *AlibabaAilabsAligenieDeviceRegisterAPIRequest {
	return poolAlibabaAilabsAligenieDeviceRegisterAPIRequest.Get().(*AlibabaAilabsAligenieDeviceRegisterAPIRequest)
}

// ReleaseAlibabaAilabsAligenieDeviceRegisterAPIRequest 将 AlibabaAilabsAligenieDeviceRegisterAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsAligenieDeviceRegisterAPIRequest(v *AlibabaAilabsAligenieDeviceRegisterAPIRequest) {
	v.Reset()
	poolAlibabaAilabsAligenieDeviceRegisterAPIRequest.Put(v)
}
