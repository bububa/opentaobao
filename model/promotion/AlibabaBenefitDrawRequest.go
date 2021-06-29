package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖接口 API请求
alibaba.benefit.draw

功能：抽奖功能，供小程序抽奖调用
业务逻辑：程序中通过奖池编号ename,业务方身份appName来查询奖池，根据授权用户(买家)来确认抽奖用户。然后程序进行抽奖流程。
小程。
安全保障：为保证数据不会越权，需要买家授，并且验证系统参数appKey。只有通过授权的，并且
appkey验证通过的，才会进入抽奖流程，否则直接失败。
因为appkey是系统参数，并且程序内部可以验证appkey和业务身份appName的关系
是否一致，所以可以保证参数appName的合法性，没有越权。
*/
type AlibabaBenefitDrawRequest struct {
    model.Params
    // 奖池唯一标识，奖池创建时即生成
    ename   string
    // 调用方AppName：规定为promotioncenter-${appId}
    appName   string
    // 调用方应用ip，非必填
    ip   string
}

// 初始化AlibabaBenefitDrawRequest对象
func NewAlibabaBenefitDrawRequest() *AlibabaBenefitDrawRequest{
    return &AlibabaBenefitDrawRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBenefitDrawRequest) GetApiMethodName() string {
    return "alibaba.benefit.draw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBenefitDrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ename Setter
// 奖池唯一标识，奖池创建时即生成
func (r *AlibabaBenefitDrawRequest) SetEname(ename string) error {
    r.ename = ename
    r.Set("ename", ename)
    return nil
}

// Ename Getter
func (r AlibabaBenefitDrawRequest) GetEname() string {
    return r.ename
}
// AppName Setter
// 调用方AppName：规定为promotioncenter-${appId}
func (r *AlibabaBenefitDrawRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

// AppName Getter
func (r AlibabaBenefitDrawRequest) GetAppName() string {
    return r.appName
}
// Ip Setter
// 调用方应用ip，非必填
func (r *AlibabaBenefitDrawRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

// Ip Getter
func (r AlibabaBenefitDrawRequest) GetIp() string {
    return r.ip
}
