package tmallcar

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarLeaseQueryloanplansAPIRequest) Reset() {
	r._loanarno = ""
	r._iproleid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseQueryloanplansAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.queryloanplans"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseQueryloanplansAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarLeaseQueryloanplansAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallCarLeaseQueryloanplansAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarLeaseQueryloanplansRequest()
	},
}

// GetTmallCarLeaseQueryloanplansRequest 从 sync.Pool 获取 TmallCarLeaseQueryloanplansAPIRequest
func GetTmallCarLeaseQueryloanplansAPIRequest() *TmallCarLeaseQueryloanplansAPIRequest {
	return poolTmallCarLeaseQueryloanplansAPIRequest.Get().(*TmallCarLeaseQueryloanplansAPIRequest)
}

// ReleaseTmallCarLeaseQueryloanplansAPIRequest 将 TmallCarLeaseQueryloanplansAPIRequest 放入 sync.Pool
func ReleaseTmallCarLeaseQueryloanplansAPIRequest(v *TmallCarLeaseQueryloanplansAPIRequest) {
	v.Reset()
	poolTmallCarLeaseQueryloanplansAPIRequest.Put(v)
}
