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
    loanarno   string
    // 客户的角色编号
    iproleid   string
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
func (r *TmallCarLeaseQueryloanplansRequest) SetLoanarno(loanarno string) error {
    r.loanarno = loanarno
    r.Set("loanarno", loanarno)
    return nil
}

// Loanarno Getter
func (r TmallCarLeaseQueryloanplansRequest) GetLoanarno() string {
    return r.loanarno
}
// Iproleid Setter
// 客户的角色编号
func (r *TmallCarLeaseQueryloanplansRequest) SetIproleid(iproleid string) error {
    r.iproleid = iproleid
    r.Set("iproleid", iproleid)
    return nil
}

// Iproleid Getter
func (r TmallCarLeaseQueryloanplansRequest) GetIproleid() string {
    return r.iproleid
}
