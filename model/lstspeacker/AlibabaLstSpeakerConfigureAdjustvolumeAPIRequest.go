package lstspeacker

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest 音箱音量调节 API请求
// alibaba.lst.speaker.configure.adjustvolume
//
// 音箱音量调节
type AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 音量直
	_volume string
	// 音量类型，val:固定值, percent:百分比
	_valueType string
}

// NewAlibabaLstSpeakerConfigureAdjustvolumeRequest 初始化AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest对象
func NewAlibabaLstSpeakerConfigureAdjustvolumeRequest() *AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest {
	return &AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest) Reset() {
	r._deviceCode = ""
	r._volume = ""
	r._valueType = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.speaker.configure.adjustvolume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceCode is DeviceCode Setter
// 设备编码
func (r *AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetVolume is Volume Setter
// 音量直
func (r *AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest) SetVolume(_volume string) error {
	r._volume = _volume
	r.Set("volume", _volume)
	return nil
}

// GetVolume Volume Getter
func (r AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest) GetVolume() string {
	return r._volume
}

// SetValueType is ValueType Setter
// 音量类型，val:固定值, percent:百分比
func (r *AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest) SetValueType(_valueType string) error {
	r._valueType = _valueType
	r.Set("value_type", _valueType)
	return nil
}

// GetValueType ValueType Getter
func (r AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest) GetValueType() string {
	return r._valueType
}

var poolAlibabaLstSpeakerConfigureAdjustvolumeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstSpeakerConfigureAdjustvolumeRequest()
	},
}

// GetAlibabaLstSpeakerConfigureAdjustvolumeRequest 从 sync.Pool 获取 AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest
func GetAlibabaLstSpeakerConfigureAdjustvolumeAPIRequest() *AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest {
	return poolAlibabaLstSpeakerConfigureAdjustvolumeAPIRequest.Get().(*AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest)
}

// ReleaseAlibabaLstSpeakerConfigureAdjustvolumeAPIRequest 将 AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstSpeakerConfigureAdjustvolumeAPIRequest(v *AlibabaLstSpeakerConfigureAdjustvolumeAPIRequest) {
	v.Reset()
	poolAlibabaLstSpeakerConfigureAdjustvolumeAPIRequest.Put(v)
}
