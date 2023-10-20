package lstspeacker

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerConfigureAPIRequest 零售通音箱配置通用泛化调用接口 API请求
// alibaba.lst.speaker.configure
//
// 零售通音箱配置通用泛化调用接口，包括内容、音量、音频等内容
type AlibabaLstSpeakerConfigureAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 命令类型setPayTime,adjustVolume，syncAudio，syncAudioAdvert
	_command string
	// 数据体，根据命令不同而不同
	_params string
}

// NewAlibabaLstSpeakerConfigureRequest 初始化AlibabaLstSpeakerConfigureAPIRequest对象
func NewAlibabaLstSpeakerConfigureRequest() *AlibabaLstSpeakerConfigureAPIRequest {
	return &AlibabaLstSpeakerConfigureAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstSpeakerConfigureAPIRequest) Reset() {
	r._deviceCode = ""
	r._command = ""
	r._params = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerConfigureAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.speaker.configure"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerConfigureAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstSpeakerConfigureAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceCode is DeviceCode Setter
// 设备编码
func (r *AlibabaLstSpeakerConfigureAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r AlibabaLstSpeakerConfigureAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetCommand is Command Setter
// 命令类型setPayTime,adjustVolume，syncAudio，syncAudioAdvert
func (r *AlibabaLstSpeakerConfigureAPIRequest) SetCommand(_command string) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// GetCommand Command Getter
func (r AlibabaLstSpeakerConfigureAPIRequest) GetCommand() string {
	return r._command
}

// SetParams is Params Setter
// 数据体，根据命令不同而不同
func (r *AlibabaLstSpeakerConfigureAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabaLstSpeakerConfigureAPIRequest) GetParams() string {
	return r._params
}

var poolAlibabaLstSpeakerConfigureAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstSpeakerConfigureRequest()
	},
}

// GetAlibabaLstSpeakerConfigureRequest 从 sync.Pool 获取 AlibabaLstSpeakerConfigureAPIRequest
func GetAlibabaLstSpeakerConfigureAPIRequest() *AlibabaLstSpeakerConfigureAPIRequest {
	return poolAlibabaLstSpeakerConfigureAPIRequest.Get().(*AlibabaLstSpeakerConfigureAPIRequest)
}

// ReleaseAlibabaLstSpeakerConfigureAPIRequest 将 AlibabaLstSpeakerConfigureAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstSpeakerConfigureAPIRequest(v *AlibabaLstSpeakerConfigureAPIRequest) {
	v.Reset()
	poolAlibabaLstSpeakerConfigureAPIRequest.Put(v)
}
