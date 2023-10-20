package guoguo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoguoguocpnborderfrontrupdateuserAPIRequest 小件员信息变更 API请求
// cainiao.guoguo.cp.nborderfrontr.updateuser
//
// 小件员信息变更
type CainiaoguoguocpnborderfrontrupdateuserAPIRequest struct {
	model.Params
	// 姓名
	_name string
	// 网点站点信息
	_workStationName string
	// 小件员员工编号
	_cpUserId string
	// 支付宝账号
	_alipayAccount string
	// 城市
	_cityName string
	// 城市行政区域编码
	_cityCode string
	// 网点站点编码
	_workStationCode string
	// 小件员所在公司编号
	_cpCode string
	// 手机号
	_mobile string
}

// NewCainiaoguoguocpnborderfrontrupdateuserRequest 初始化CainiaoguoguocpnborderfrontrupdateuserAPIRequest对象
func NewCainiaoguoguocpnborderfrontrupdateuserRequest() *CainiaoguoguocpnborderfrontrupdateuserAPIRequest {
	return &CainiaoguoguocpnborderfrontrupdateuserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoguoguocpnborderfrontrupdateuserAPIRequest) GetApiMethodName() string {
	return "cainiao.guoguo.cp.nborderfrontr.updateuser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoguoguocpnborderfrontrupdateuserAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoguoguocpnborderfrontrupdateuserAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 姓名
func (r *CainiaoguoguocpnborderfrontrupdateuserAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r CainiaoguoguocpnborderfrontrupdateuserAPIRequest) GetName() string {
	return r._name
}

// SetWorkStationName is WorkStationName Setter
// 网点站点信息
func (r *CainiaoguoguocpnborderfrontrupdateuserAPIRequest) SetWorkStationName(_workStationName string) error {
	r._workStationName = _workStationName
	r.Set("work_station_name", _workStationName)
	return nil
}

// GetWorkStationName WorkStationName Getter
func (r CainiaoguoguocpnborderfrontrupdateuserAPIRequest) GetWorkStationName() string {
	return r._workStationName
}

// SetCpUserId is CpUserId Setter
// 小件员员工编号
func (r *CainiaoguoguocpnborderfrontrupdateuserAPIRequest) SetCpUserId(_cpUserId string) error {
	r._cpUserId = _cpUserId
	r.Set("cp_user_id", _cpUserId)
	return nil
}

// GetCpUserId CpUserId Getter
func (r CainiaoguoguocpnborderfrontrupdateuserAPIRequest) GetCpUserId() string {
	return r._cpUserId
}

// SetAlipayAccount is AlipayAccount Setter
// 支付宝账号
func (r *CainiaoguoguocpnborderfrontrupdateuserAPIRequest) SetAlipayAccount(_alipayAccount string) error {
	r._alipayAccount = _alipayAccount
	r.Set("alipay_account", _alipayAccount)
	return nil
}

// GetAlipayAccount AlipayAccount Getter
func (r CainiaoguoguocpnborderfrontrupdateuserAPIRequest) GetAlipayAccount() string {
	return r._alipayAccount
}

// SetCityName is CityName Setter
// 城市
func (r *CainiaoguoguocpnborderfrontrupdateuserAPIRequest) SetCityName(_cityName string) error {
	r._cityName = _cityName
	r.Set("city_name", _cityName)
	return nil
}

// GetCityName CityName Getter
func (r CainiaoguoguocpnborderfrontrupdateuserAPIRequest) GetCityName() string {
	return r._cityName
}

// SetCityCode is CityCode Setter
// 城市行政区域编码
func (r *CainiaoguoguocpnborderfrontrupdateuserAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r CainiaoguoguocpnborderfrontrupdateuserAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetWorkStationCode is WorkStationCode Setter
// 网点站点编码
func (r *CainiaoguoguocpnborderfrontrupdateuserAPIRequest) SetWorkStationCode(_workStationCode string) error {
	r._workStationCode = _workStationCode
	r.Set("work_station_code", _workStationCode)
	return nil
}

// GetWorkStationCode WorkStationCode Getter
func (r CainiaoguoguocpnborderfrontrupdateuserAPIRequest) GetWorkStationCode() string {
	return r._workStationCode
}

// SetCpCode is CpCode Setter
// 小件员所在公司编号
func (r *CainiaoguoguocpnborderfrontrupdateuserAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaoguoguocpnborderfrontrupdateuserAPIRequest) GetCpCode() string {
	return r._cpCode
}

// SetMobile is Mobile Setter
// 手机号
func (r *CainiaoguoguocpnborderfrontrupdateuserAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r CainiaoguoguocpnborderfrontrupdateuserAPIRequest) GetMobile() string {
	return r._mobile
}
