package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstspeakerconfigureAPIRequest 零售通音箱配置通用泛化调用接口 API请求
// alibaba.lst.speaker.configure
//
// 零售通音箱配置通用泛化调用接口，包括内容、音量、音频等内容
type AlibabalstspeakerconfigureAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 命令类型setPayTime,adjustVolume，syncAudio，syncAudioAdvert
	_command string
	// 数据体，根据命令不同而不同
	_params string
}

// NewAlibabalstspeakerconfigureRequest 初始化AlibabalstspeakerconfigureAPIRequest对象
func NewAlibabalstspeakerconfigureRequest() *AlibabalstspeakerconfigureAPIRequest {
	return &AlibabalstspeakerconfigureAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstspeakerconfigureAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.speaker.configure"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstspeakerconfigureAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstspeakerconfigureAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceCode is DeviceCode Setter
// 设备编码
func (r *AlibabalstspeakerconfigureAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r AlibabalstspeakerconfigureAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetCommand is Command Setter
// 命令类型setPayTime,adjustVolume，syncAudio，syncAudioAdvert
func (r *AlibabalstspeakerconfigureAPIRequest) SetCommand(_command string) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// GetCommand Command Getter
func (r AlibabalstspeakerconfigureAPIRequest) GetCommand() string {
	return r._command
}

// SetParams is Params Setter
// 数据体，根据命令不同而不同
func (r *AlibabalstspeakerconfigureAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabalstspeakerconfigureAPIRequest) GetParams() string {
	return r._params
}
