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
    ename   string
    // 调用方AppName：规定为promotioncenter-${appId}
    appName   string
    // 调用方应用ip，非必填
    ip   string
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
func (r *AlibabaBeneiftDrawRequest) SetEname(ename string) error {
    r.ename = ename
    r.Set("ename", ename)
    return nil
}

// Ename Getter
func (r AlibabaBeneiftDrawRequest) GetEname() string {
    return r.ename
}
// AppName Setter
// 调用方AppName：规定为promotioncenter-${appId}
func (r *AlibabaBeneiftDrawRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

// AppName Getter
func (r AlibabaBeneiftDrawRequest) GetAppName() string {
    return r.appName
}
// Ip Setter
// 调用方应用ip，非必填
func (r *AlibabaBeneiftDrawRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

// Ip Getter
func (r AlibabaBeneiftDrawRequest) GetIp() string {
    return r.ip
}
