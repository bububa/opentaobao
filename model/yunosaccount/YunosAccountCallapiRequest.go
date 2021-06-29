package yunosaccount

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
调用YunOS账号开放API APIRequest
yunos.account.callapi

YunOS账号客户端对外开放的api通过top暴露
*/
type YunosAccountCallapiRequest struct {
    model.Params

    // API版本号
    version   string 

    // 调用的API名称
    api   string 

    // 时间戳，精确到秒；账号服务端会校验该值与服务器当前时间戳的差值，超过一定范围则拒绝请求
    timeStamp   string 

    // 业务参数
    params   string 

    // 应用签名的MD5值
    authSign   string 

}

func NewYunosAccountCallapiRequest() *YunosAccountCallapiRequest{
    return &YunosAccountCallapiRequest{
        Params: model.NewParams(),
    }
}

func (r YunosAccountCallapiRequest) GetApiMethodName() string {
    return "yunos.account.callapi"
}

func (r YunosAccountCallapiRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosAccountCallapiRequest) SetVersion(version string) error {
    r.version = version
    r.Set("version", version)
    return nil
}

func (r YunosAccountCallapiRequest) GetVersion() string {
    return r.version
}

func (r *YunosAccountCallapiRequest) SetApi(api string) error {
    r.api = api
    r.Set("api", api)
    return nil
}

func (r YunosAccountCallapiRequest) GetApi() string {
    return r.api
}

func (r *YunosAccountCallapiRequest) SetTimeStamp(timeStamp string) error {
    r.timeStamp = timeStamp
    r.Set("time_stamp", timeStamp)
    return nil
}

func (r YunosAccountCallapiRequest) GetTimeStamp() string {
    return r.timeStamp
}

func (r *YunosAccountCallapiRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r YunosAccountCallapiRequest) GetParams() string {
    return r.params
}

func (r *YunosAccountCallapiRequest) SetAuthSign(authSign string) error {
    r.authSign = authSign
    r.Set("auth_sign", authSign)
    return nil
}

func (r YunosAccountCallapiRequest) GetAuthSign() string {
    return r.authSign
}

