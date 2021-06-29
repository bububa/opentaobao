package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售云小程序获取登录用户信息 API请求
alibaba.lsy.miniapp.user.get

零售云小程序，通过授权码获取登录的卖家账号信息
*/
type AlibabaLsyMiniappUserGetRequest struct {
    model.Params
    // 当前时间戳，毫秒
    _timeStamp   string
    // 获取用户信息的授权码，在小程序中获取
    _code   string
    // 请求参数签名，sha1(所有入参+appSecret，按字符串升序排列)
    _signature   string
    // 系统分配的小程序ID
    _appId   string
}

// 初始化AlibabaLsyMiniappUserGetRequest对象
func NewAlibabaLsyMiniappUserGetRequest() *AlibabaLsyMiniappUserGetRequest{
    return &AlibabaLsyMiniappUserGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyMiniappUserGetRequest) GetApiMethodName() string {
    return "alibaba.lsy.miniapp.user.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyMiniappUserGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TimeStamp Setter
// 当前时间戳，毫秒
func (r *AlibabaLsyMiniappUserGetRequest) SetTimeStamp(_timeStamp string) error {
    r._timeStamp = _timeStamp
    r.Set("time_stamp", _timeStamp)
    return nil
}

// TimeStamp Getter
func (r AlibabaLsyMiniappUserGetRequest) GetTimeStamp() string {
    return r._timeStamp
}
// Code Setter
// 获取用户信息的授权码，在小程序中获取
func (r *AlibabaLsyMiniappUserGetRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaLsyMiniappUserGetRequest) GetCode() string {
    return r._code
}
// Signature Setter
// 请求参数签名，sha1(所有入参+appSecret，按字符串升序排列)
func (r *AlibabaLsyMiniappUserGetRequest) SetSignature(_signature string) error {
    r._signature = _signature
    r.Set("signature", _signature)
    return nil
}

// Signature Getter
func (r AlibabaLsyMiniappUserGetRequest) GetSignature() string {
    return r._signature
}
// AppId Setter
// 系统分配的小程序ID
func (r *AlibabaLsyMiniappUserGetRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r AlibabaLsyMiniappUserGetRequest) GetAppId() string {
    return r._appId
}
