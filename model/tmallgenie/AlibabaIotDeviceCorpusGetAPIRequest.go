package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIotDeviceCorpusGetAPIRequest IoT设备支持语料获取 API请求
// alibaba.iot.device.corpus.get
//
// ISV通过该接口获取天猫精灵IoT设备支持控制或查询的语料
type AlibabaIotDeviceCorpusGetAPIRequest struct {
	model.Params
	// 天猫精灵开放用户id
	_userOpenId string
	// 天猫精灵开放的client id
	_clientId string
	// iot设备id
	_devId string
}

// NewAlibabaIotDeviceCorpusGetRequest 初始化AlibabaIotDeviceCorpusGetAPIRequest对象
func NewAlibabaIotDeviceCorpusGetRequest() *AlibabaIotDeviceCorpusGetAPIRequest {
	return &AlibabaIotDeviceCorpusGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIotDeviceCorpusGetAPIRequest) Reset() {
	r._userOpenId = ""
	r._clientId = ""
	r._devId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIotDeviceCorpusGetAPIRequest) GetApiMethodName() string {
	return "alibaba.iot.device.corpus.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIotDeviceCorpusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIotDeviceCorpusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserOpenId is UserOpenId Setter
// 天猫精灵开放用户id
func (r *AlibabaIotDeviceCorpusGetAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// GetUserOpenId UserOpenId Getter
func (r AlibabaIotDeviceCorpusGetAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// SetClientId is ClientId Setter
// 天猫精灵开放的client id
func (r *AlibabaIotDeviceCorpusGetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaIotDeviceCorpusGetAPIRequest) GetClientId() string {
	return r._clientId
}

// SetDevId is DevId Setter
// iot设备id
func (r *AlibabaIotDeviceCorpusGetAPIRequest) SetDevId(_devId string) error {
	r._devId = _devId
	r.Set("dev_id", _devId)
	return nil
}

// GetDevId DevId Getter
func (r AlibabaIotDeviceCorpusGetAPIRequest) GetDevId() string {
	return r._devId
}

var poolAlibabaIotDeviceCorpusGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIotDeviceCorpusGetRequest()
	},
}

// GetAlibabaIotDeviceCorpusGetRequest 从 sync.Pool 获取 AlibabaIotDeviceCorpusGetAPIRequest
func GetAlibabaIotDeviceCorpusGetAPIRequest() *AlibabaIotDeviceCorpusGetAPIRequest {
	return poolAlibabaIotDeviceCorpusGetAPIRequest.Get().(*AlibabaIotDeviceCorpusGetAPIRequest)
}

// ReleaseAlibabaIotDeviceCorpusGetAPIRequest 将 AlibabaIotDeviceCorpusGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIotDeviceCorpusGetAPIRequest(v *AlibabaIotDeviceCorpusGetAPIRequest) {
	v.Reset()
	poolAlibabaIotDeviceCorpusGetAPIRequest.Put(v)
}
