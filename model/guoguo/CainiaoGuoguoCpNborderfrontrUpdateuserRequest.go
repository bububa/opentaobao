package guoguo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小件员信息变更 API请求
cainiao.guoguo.cp.nborderfrontr.updateuser

小件员信息变更
*/
type CainiaoGuoguoCpNborderfrontrUpdateuserRequest struct {
    model.Params
    // 姓名
    name   string
    // 网点站点信息
    workStationName   string
    // 小件员员工编号
    cpUserId   string
    // 支付宝账号
    alipayAccount   string
    // 城市
    cityName   string
    // 城市行政区域编码
    cityCode   string
    // 网点站点编码
    workStationCode   string
    // 小件员所在公司编号
    cpCode   string
    // 手机号
    mobile   string
}

// 初始化CainiaoGuoguoCpNborderfrontrUpdateuserRequest对象
func NewCainiaoGuoguoCpNborderfrontrUpdateuserRequest() *CainiaoGuoguoCpNborderfrontrUpdateuserRequest{
    return &CainiaoGuoguoCpNborderfrontrUpdateuserRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetApiMethodName() string {
    return "cainiao.guoguo.cp.nborderfrontr.updateuser"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 姓名
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetName() string {
    return r.name
}
// WorkStationName Setter
// 网点站点信息
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetWorkStationName(workStationName string) error {
    r.workStationName = workStationName
    r.Set("work_station_name", workStationName)
    return nil
}

// WorkStationName Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetWorkStationName() string {
    return r.workStationName
}
// CpUserId Setter
// 小件员员工编号
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetCpUserId(cpUserId string) error {
    r.cpUserId = cpUserId
    r.Set("cp_user_id", cpUserId)
    return nil
}

// CpUserId Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetCpUserId() string {
    return r.cpUserId
}
// AlipayAccount Setter
// 支付宝账号
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetAlipayAccount(alipayAccount string) error {
    r.alipayAccount = alipayAccount
    r.Set("alipay_account", alipayAccount)
    return nil
}

// AlipayAccount Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetAlipayAccount() string {
    return r.alipayAccount
}
// CityName Setter
// 城市
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetCityName(cityName string) error {
    r.cityName = cityName
    r.Set("city_name", cityName)
    return nil
}

// CityName Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetCityName() string {
    return r.cityName
}
// CityCode Setter
// 城市行政区域编码
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetCityCode() string {
    return r.cityCode
}
// WorkStationCode Setter
// 网点站点编码
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetWorkStationCode(workStationCode string) error {
    r.workStationCode = workStationCode
    r.Set("work_station_code", workStationCode)
    return nil
}

// WorkStationCode Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetWorkStationCode() string {
    return r.workStationCode
}
// CpCode Setter
// 小件员所在公司编号
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetCpCode() string {
    return r.cpCode
}
// Mobile Setter
// 手机号
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetMobile() string {
    return r.mobile
}
