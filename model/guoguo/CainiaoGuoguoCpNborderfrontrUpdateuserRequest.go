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
    _name   string
    // 网点站点信息
    _workStationName   string
    // 小件员员工编号
    _cpUserId   string
    // 支付宝账号
    _alipayAccount   string
    // 城市
    _cityName   string
    // 城市行政区域编码
    _cityCode   string
    // 网点站点编码
    _workStationCode   string
    // 小件员所在公司编号
    _cpCode   string
    // 手机号
    _mobile   string
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
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetName() string {
    return r._name
}
// WorkStationName Setter
// 网点站点信息
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetWorkStationName(_workStationName string) error {
    r._workStationName = _workStationName
    r.Set("work_station_name", _workStationName)
    return nil
}

// WorkStationName Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetWorkStationName() string {
    return r._workStationName
}
// CpUserId Setter
// 小件员员工编号
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetCpUserId(_cpUserId string) error {
    r._cpUserId = _cpUserId
    r.Set("cp_user_id", _cpUserId)
    return nil
}

// CpUserId Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetCpUserId() string {
    return r._cpUserId
}
// AlipayAccount Setter
// 支付宝账号
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetAlipayAccount(_alipayAccount string) error {
    r._alipayAccount = _alipayAccount
    r.Set("alipay_account", _alipayAccount)
    return nil
}

// AlipayAccount Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetAlipayAccount() string {
    return r._alipayAccount
}
// CityName Setter
// 城市
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetCityName(_cityName string) error {
    r._cityName = _cityName
    r.Set("city_name", _cityName)
    return nil
}

// CityName Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetCityName() string {
    return r._cityName
}
// CityCode Setter
// 城市行政区域编码
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetCityCode() string {
    return r._cityCode
}
// WorkStationCode Setter
// 网点站点编码
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetWorkStationCode(_workStationCode string) error {
    r._workStationCode = _workStationCode
    r.Set("work_station_code", _workStationCode)
    return nil
}

// WorkStationCode Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetWorkStationCode() string {
    return r._workStationCode
}
// CpCode Setter
// 小件员所在公司编号
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetCpCode(_cpCode string) error {
    r._cpCode = _cpCode
    r.Set("cp_code", _cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetCpCode() string {
    return r._cpCode
}
// Mobile Setter
// 手机号
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserRequest) GetMobile() string {
    return r._mobile
}
