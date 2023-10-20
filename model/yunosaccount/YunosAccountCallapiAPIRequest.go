package yunosaccount

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosaccountcallapiAPIRequest 调用YunOS账号开放API API请求
// yunos.account.callapi
//
// YunOS账号客户端对外开放的api通过top暴露
type YunosaccountcallapiAPIRequest struct {
	model.Params
	// API版本号
	_version string
	// 调用的API名称
	_api string
	// 时间戳，精确到秒；账号服务端会校验该值与服务器当前时间戳的差值，超过一定范围则拒绝请求
	_timeStamp string
	// 业务参数
	_params string
	// 应用签名的MD5值
	_authSign string
}

// NewYunosaccountcallapiRequest 初始化YunosaccountcallapiAPIRequest对象
func NewYunosaccountcallapiRequest() *YunosaccountcallapiAPIRequest {
	return &YunosaccountcallapiAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosaccountcallapiAPIRequest) GetApiMethodName() string {
	return "yunos.account.callapi"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosaccountcallapiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosaccountcallapiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVersion is Version Setter
// API版本号
func (r *YunosaccountcallapiAPIRequest) SetVersion(_version string) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r YunosaccountcallapiAPIRequest) GetVersion() string {
	return r._version
}

// SetApi is Api Setter
// 调用的API名称
func (r *YunosaccountcallapiAPIRequest) SetApi(_api string) error {
	r._api = _api
	r.Set("api", _api)
	return nil
}

// GetApi Api Getter
func (r YunosaccountcallapiAPIRequest) GetApi() string {
	return r._api
}

// SetTimeStamp is TimeStamp Setter
// 时间戳，精确到秒；账号服务端会校验该值与服务器当前时间戳的差值，超过一定范围则拒绝请求
func (r *YunosaccountcallapiAPIRequest) SetTimeStamp(_timeStamp string) error {
	r._timeStamp = _timeStamp
	r.Set("time_stamp", _timeStamp)
	return nil
}

// GetTimeStamp TimeStamp Getter
func (r YunosaccountcallapiAPIRequest) GetTimeStamp() string {
	return r._timeStamp
}

// SetParams is Params Setter
// 业务参数
func (r *YunosaccountcallapiAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r YunosaccountcallapiAPIRequest) GetParams() string {
	return r._params
}

// SetAuthSign is AuthSign Setter
// 应用签名的MD5值
func (r *YunosaccountcallapiAPIRequest) SetAuthSign(_authSign string) error {
	r._authSign = _authSign
	r.Set("auth_sign", _authSign)
	return nil
}

// GetAuthSign AuthSign Getter
func (r YunosaccountcallapiAPIRequest) GetAuthSign() string {
	return r._authSign
}
