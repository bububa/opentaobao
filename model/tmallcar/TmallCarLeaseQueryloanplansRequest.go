package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后查询还款计划 API请求
tmall.car.lease.queryloanplans

天猫开新车租后查询还款计划
*/
type TmallCarLeaseQueryloanplansRequest struct {
    model.Params
    // 合约编号
    _loanarno   string
    // 客户的角色编号
    _iproleid   string
}

// 初始化TmallCarLeaseQueryloanplansRequest对象
func NewTmallCarLeaseQueryloanplansRequest() *TmallCarLeaseQueryloanplansRequest{
    return &TmallCarLeaseQueryloanplansRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseQueryloanplansRequest) GetApiMethodName() string {
    return "tmall.car.lease.queryloanplans"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseQueryloanplansRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Loanarno Setter
// 合约编号
func (r *TmallCarLeaseQueryloanplansRequest) SetLoanarno(_loanarno string) error {
    r._loanarno = _loanarno
    r.Set("loanarno", _loanarno)
    return nil
}

// Loanarno Getter
func (r TmallCarLeaseQueryloanplansRequest) GetLoanarno() string {
    return r._loanarno
}
// Iproleid Setter
// 客户的角色编号
func (r *TmallCarLeaseQueryloanplansRequest) SetIproleid(_iproleid string) error {
    r._iproleid = _iproleid
    r.Set("iproleid", _iproleid)
    return nil
}

// Iproleid Getter
func (r TmallCarLeaseQueryloanplansRequest) GetIproleid() string {
    return r._iproleid
}
