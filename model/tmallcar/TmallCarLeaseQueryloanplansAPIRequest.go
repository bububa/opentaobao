package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseQueryloanplansAPIRequest 天猫开新车租后查询还款计划 API请求
// tmall.car.lease.queryloanplans
//
// 天猫开新车租后查询还款计划
type TmallCarLeaseQueryloanplansAPIRequest struct {
	model.Params
	// 合约编号
	_loanarno string
	// 客户的角色编号
	_iproleid string
}

// NewTmallCarLeaseQueryloanplansRequest 初始化TmallCarLeaseQueryloanplansAPIRequest对象
func NewTmallCarLeaseQueryloanplansRequest() *TmallCarLeaseQueryloanplansAPIRequest {
	return &TmallCarLeaseQueryloanplansAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseQueryloanplansAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.queryloanplans"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseQueryloanplansAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLoanarno is Loanarno Setter
// 合约编号
func (r *TmallCarLeaseQueryloanplansAPIRequest) SetLoanarno(_loanarno string) error {
	r._loanarno = _loanarno
	r.Set("loanarno", _loanarno)
	return nil
}

// GetLoanarno Loanarno Getter
func (r TmallCarLeaseQueryloanplansAPIRequest) GetLoanarno() string {
	return r._loanarno
}

// SetIproleid is Iproleid Setter
// 客户的角色编号
func (r *TmallCarLeaseQueryloanplansAPIRequest) SetIproleid(_iproleid string) error {
	r._iproleid = _iproleid
	r.Set("iproleid", _iproleid)
	return nil
}

// GetIproleid Iproleid Getter
func (r TmallCarLeaseQueryloanplansAPIRequest) GetIproleid() string {
	return r._iproleid
}
