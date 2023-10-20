package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleasequeryloanplansAPIRequest 天猫开新车租后查询还款计划 API请求
// tmall.car.lease.queryloanplans
//
// 天猫开新车租后查询还款计划
type TmallcarleasequeryloanplansAPIRequest struct {
	model.Params
	// 合约编号
	_loanarno string
	// 客户的角色编号
	_iproleid string
}

// NewTmallcarleasequeryloanplansRequest 初始化TmallcarleasequeryloanplansAPIRequest对象
func NewTmallcarleasequeryloanplansRequest() *TmallcarleasequeryloanplansAPIRequest {
	return &TmallcarleasequeryloanplansAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarleasequeryloanplansAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.queryloanplans"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarleasequeryloanplansAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarleasequeryloanplansAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLoanarno is Loanarno Setter
// 合约编号
func (r *TmallcarleasequeryloanplansAPIRequest) SetLoanarno(_loanarno string) error {
	r._loanarno = _loanarno
	r.Set("loanarno", _loanarno)
	return nil
}

// GetLoanarno Loanarno Getter
func (r TmallcarleasequeryloanplansAPIRequest) GetLoanarno() string {
	return r._loanarno
}

// SetIproleid is Iproleid Setter
// 客户的角色编号
func (r *TmallcarleasequeryloanplansAPIRequest) SetIproleid(_iproleid string) error {
	r._iproleid = _iproleid
	r.Set("iproleid", _iproleid)
	return nil
}

// GetIproleid Iproleid Getter
func (r TmallcarleasequeryloanplansAPIRequest) GetIproleid() string {
	return r._iproleid
}
