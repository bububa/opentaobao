package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖接口 API请求
alibaba.beneift.draw

抽奖接口
*/
type AlibabaBeneiftDrawRequest struct {
    model.Params
    // 奖池唯一标识，奖池创建时即生成
    _ename   string
    // 调用方AppName：规定为promotioncenter-${appId}
    _appName   string
    // 调用方应用ip，非必填
    _ip   string
}

// 初始化AlibabaBeneiftDrawRequest对象
func NewAlibabaBeneiftDrawRequest() *AlibabaBeneiftDrawRequest{
    return &AlibabaBeneiftDrawRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBeneiftDrawRequest) GetApiMethodName() string {
    return "alibaba.beneift.draw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBeneiftDrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ename Setter
// 奖池唯一标识，奖池创建时即生成
func (r *AlibabaBeneiftDrawRequest) SetEname(_ename string) error {
    r._ename = _ename
    r.Set("ename", _ename)
    return nil
}

// Ename Getter
func (r AlibabaBeneiftDrawRequest) GetEname() string {
    return r._ename
}
// AppName Setter
// 调用方AppName：规定为promotioncenter-${appId}
func (r *AlibabaBeneiftDrawRequest) SetAppName(_appName string) error {
    r._appName = _appName
    r.Set("app_name", _appName)
    return nil
}

// AppName Getter
func (r AlibabaBeneiftDrawRequest) GetAppName() string {
    return r._appName
}
// Ip Setter
// 调用方应用ip，非必填
func (r *AlibabaBeneiftDrawRequest) SetIp(_ip string) error {
    r._ip = _ip
    r.Set("ip", _ip)
    return nil
}

// Ip Getter
func (r AlibabaBeneiftDrawRequest) GetIp() string {
    return r._ip
}
