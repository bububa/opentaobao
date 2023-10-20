package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaiotdevicecorpusgetAPIRequest IoT设备支持语料获取 API请求
// alibaba.iot.device.corpus.get
//
// ISV通过该接口获取天猫精灵IoT设备支持控制或查询的语料
type AlibabaiotdevicecorpusgetAPIRequest struct {
	model.Params
	// 天猫精灵开放用户id
	_userOpenId string
	// 天猫精灵开放的client id
	_clientId string
	// iot设备id
	_devId string
}

// NewAlibabaiotdevicecorpusgetRequest 初始化AlibabaiotdevicecorpusgetAPIRequest对象
func NewAlibabaiotdevicecorpusgetRequest() *AlibabaiotdevicecorpusgetAPIRequest {
	return &AlibabaiotdevicecorpusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaiotdevicecorpusgetAPIRequest) GetApiMethodName() string {
	return "alibaba.iot.device.corpus.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaiotdevicecorpusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaiotdevicecorpusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserOpenId is UserOpenId Setter
// 天猫精灵开放用户id
func (r *AlibabaiotdevicecorpusgetAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// GetUserOpenId UserOpenId Getter
func (r AlibabaiotdevicecorpusgetAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// SetClientId is ClientId Setter
// 天猫精灵开放的client id
func (r *AlibabaiotdevicecorpusgetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaiotdevicecorpusgetAPIRequest) GetClientId() string {
	return r._clientId
}

// SetDevId is DevId Setter
// iot设备id
func (r *AlibabaiotdevicecorpusgetAPIRequest) SetDevId(_devId string) error {
	r._devId = _devId
	r.Set("dev_id", _devId)
	return nil
}

// GetDevId DevId Getter
func (r AlibabaiotdevicecorpusgetAPIRequest) GetDevId() string {
	return r._devId
}
