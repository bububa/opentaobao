package guoguo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest 小件员信息变更 API请求
// cainiao.guoguo.cp.nborderfrontr.updateuser
//
// 小件员信息变更
type CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest struct {
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

// NewCainiaoGuoguoCpNborderfrontrUpdateuserRequest 初始化CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest对象
func NewCainiaoGuoguoCpNborderfrontrUpdateuserRequest() *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest {
	return &CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest{
		Params: model.NewParams(9),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) Reset() {
	r._name = ""
	r._workStationName = ""
	r._cpUserId = ""
	r._alipayAccount = ""
	r._cityName = ""
	r._cityCode = ""
	r._workStationCode = ""
	r._cpCode = ""
	r._mobile = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) GetApiMethodName() string {
	return "cainiao.guoguo.cp.nborderfrontr.updateuser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 姓名
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) GetName() string {
	return r._name
}

// SetWorkStationName is WorkStationName Setter
// 网点站点信息
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) SetWorkStationName(_workStationName string) error {
	r._workStationName = _workStationName
	r.Set("work_station_name", _workStationName)
	return nil
}

// GetWorkStationName WorkStationName Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) GetWorkStationName() string {
	return r._workStationName
}

// SetCpUserId is CpUserId Setter
// 小件员员工编号
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) SetCpUserId(_cpUserId string) error {
	r._cpUserId = _cpUserId
	r.Set("cp_user_id", _cpUserId)
	return nil
}

// GetCpUserId CpUserId Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) GetCpUserId() string {
	return r._cpUserId
}

// SetAlipayAccount is AlipayAccount Setter
// 支付宝账号
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) SetAlipayAccount(_alipayAccount string) error {
	r._alipayAccount = _alipayAccount
	r.Set("alipay_account", _alipayAccount)
	return nil
}

// GetAlipayAccount AlipayAccount Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) GetAlipayAccount() string {
	return r._alipayAccount
}

// SetCityName is CityName Setter
// 城市
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) SetCityName(_cityName string) error {
	r._cityName = _cityName
	r.Set("city_name", _cityName)
	return nil
}

// GetCityName CityName Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) GetCityName() string {
	return r._cityName
}

// SetCityCode is CityCode Setter
// 城市行政区域编码
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetWorkStationCode is WorkStationCode Setter
// 网点站点编码
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) SetWorkStationCode(_workStationCode string) error {
	r._workStationCode = _workStationCode
	r.Set("work_station_code", _workStationCode)
	return nil
}

// GetWorkStationCode WorkStationCode Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) GetWorkStationCode() string {
	return r._workStationCode
}

// SetCpCode is CpCode Setter
// 小件员所在公司编号
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) GetCpCode() string {
	return r._cpCode
}

// SetMobile is Mobile Setter
// 手机号
func (r *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) GetMobile() string {
	return r._mobile
}

var poolCainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoGuoguoCpNborderfrontrUpdateuserRequest()
	},
}

// GetCainiaoGuoguoCpNborderfrontrUpdateuserRequest 从 sync.Pool 获取 CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest
func GetCainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest() *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest {
	return poolCainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest.Get().(*CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest)
}

// ReleaseCainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest 将 CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest 放入 sync.Pool
func ReleaseCainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest(v *CainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest) {
	v.Reset()
	poolCainiaoGuoguoCpNborderfrontrUpdateuserAPIRequest.Put(v)
}
