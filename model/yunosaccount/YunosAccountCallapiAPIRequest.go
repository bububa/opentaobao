package yunosaccount

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosAccountCallapiAPIRequest 调用YunOS账号开放API API请求
// yunos.account.callapi
//
// YunOS账号客户端对外开放的api通过top暴露
type YunosAccountCallapiAPIRequest struct {
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

// NewYunosAccountCallapiRequest 初始化YunosAccountCallapiAPIRequest对象
func NewYunosAccountCallapiRequest() *YunosAccountCallapiAPIRequest {
	return &YunosAccountCallapiAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosAccountCallapiAPIRequest) Reset() {
	r._version = ""
	r._api = ""
	r._timeStamp = ""
	r._params = ""
	r._authSign = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAccountCallapiAPIRequest) GetApiMethodName() string {
	return "yunos.account.callapi"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAccountCallapiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosAccountCallapiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVersion is Version Setter
// API版本号
func (r *YunosAccountCallapiAPIRequest) SetVersion(_version string) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r YunosAccountCallapiAPIRequest) GetVersion() string {
	return r._version
}

// SetApi is Api Setter
// 调用的API名称
func (r *YunosAccountCallapiAPIRequest) SetApi(_api string) error {
	r._api = _api
	r.Set("api", _api)
	return nil
}

// GetApi Api Getter
func (r YunosAccountCallapiAPIRequest) GetApi() string {
	return r._api
}

// SetTimeStamp is TimeStamp Setter
// 时间戳，精确到秒；账号服务端会校验该值与服务器当前时间戳的差值，超过一定范围则拒绝请求
func (r *YunosAccountCallapiAPIRequest) SetTimeStamp(_timeStamp string) error {
	r._timeStamp = _timeStamp
	r.Set("time_stamp", _timeStamp)
	return nil
}

// GetTimeStamp TimeStamp Getter
func (r YunosAccountCallapiAPIRequest) GetTimeStamp() string {
	return r._timeStamp
}

// SetParams is Params Setter
// 业务参数
func (r *YunosAccountCallapiAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r YunosAccountCallapiAPIRequest) GetParams() string {
	return r._params
}

// SetAuthSign is AuthSign Setter
// 应用签名的MD5值
func (r *YunosAccountCallapiAPIRequest) SetAuthSign(_authSign string) error {
	r._authSign = _authSign
	r.Set("auth_sign", _authSign)
	return nil
}

// GetAuthSign AuthSign Getter
func (r YunosAccountCallapiAPIRequest) GetAuthSign() string {
	return r._authSign
}

var poolYunosAccountCallapiAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosAccountCallapiRequest()
	},
}

// GetYunosAccountCallapiRequest 从 sync.Pool 获取 YunosAccountCallapiAPIRequest
func GetYunosAccountCallapiAPIRequest() *YunosAccountCallapiAPIRequest {
	return poolYunosAccountCallapiAPIRequest.Get().(*YunosAccountCallapiAPIRequest)
}

// ReleaseYunosAccountCallapiAPIRequest 将 YunosAccountCallapiAPIRequest 放入 sync.Pool
func ReleaseYunosAccountCallapiAPIRequest(v *YunosAccountCallapiAPIRequest) {
	v.Reset()
	poolYunosAccountCallapiAPIRequest.Put(v)
}
