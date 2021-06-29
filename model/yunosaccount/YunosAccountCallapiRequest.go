package yunosaccount

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
调用YunOS账号开放API API请求
yunos.account.callapi

YunOS账号客户端对外开放的api通过top暴露
*/
type YunosAccountCallapiRequest struct {
    model.Params
    // API版本号
    _version   string
    // 调用的API名称
    _api   string
    // 时间戳，精确到秒；账号服务端会校验该值与服务器当前时间戳的差值，超过一定范围则拒绝请求
    _timeStamp   string
    // 业务参数
    _params   string
    // 应用签名的MD5值
    _authSign   string
}

// 初始化YunosAccountCallapiRequest对象
func NewYunosAccountCallapiRequest() *YunosAccountCallapiRequest{
    return &YunosAccountCallapiRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosAccountCallapiRequest) GetApiMethodName() string {
    return "yunos.account.callapi"
}

// IRequest interface 方法, 获取API参数
func (r YunosAccountCallapiRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Version Setter
// API版本号
func (r *YunosAccountCallapiRequest) SetVersion(_version string) error {
    r._version = _version
    r.Set("version", _version)
    return nil
}

// Version Getter
func (r YunosAccountCallapiRequest) GetVersion() string {
    return r._version
}
// Api Setter
// 调用的API名称
func (r *YunosAccountCallapiRequest) SetApi(_api string) error {
    r._api = _api
    r.Set("api", _api)
    return nil
}

// Api Getter
func (r YunosAccountCallapiRequest) GetApi() string {
    return r._api
}
// TimeStamp Setter
// 时间戳，精确到秒；账号服务端会校验该值与服务器当前时间戳的差值，超过一定范围则拒绝请求
func (r *YunosAccountCallapiRequest) SetTimeStamp(_timeStamp string) error {
    r._timeStamp = _timeStamp
    r.Set("time_stamp", _timeStamp)
    return nil
}

// TimeStamp Getter
func (r YunosAccountCallapiRequest) GetTimeStamp() string {
    return r._timeStamp
}
// Params Setter
// 业务参数
func (r *YunosAccountCallapiRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r YunosAccountCallapiRequest) GetParams() string {
    return r._params
}
// AuthSign Setter
// 应用签名的MD5值
func (r *YunosAccountCallapiRequest) SetAuthSign(_authSign string) error {
    r._authSign = _authSign
    r.Set("auth_sign", _authSign)
    return nil
}

// AuthSign Getter
func (r YunosAccountCallapiRequest) GetAuthSign() string {
    return r._authSign
}
