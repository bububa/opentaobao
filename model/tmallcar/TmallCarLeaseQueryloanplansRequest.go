package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后查询还款计划 APIRequest
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

func NewTmallCarLeaseQueryloanplansRequest() *TmallCarLeaseQueryloanplansRequest{
    return &TmallCarLeaseQueryloanplansRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarLeaseQueryloanplansRequest) GetApiMethodName() string {
    return "tmall.car.lease.queryloanplans"
}

func (r TmallCarLeaseQueryloanplansRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarLeaseQueryloanplansRequest) SetLoanarno(loanarno string) error {
    r.loanarno = loanarno
    r.Set("loanarno", loanarno)
    return nil
}

func (r TmallCarLeaseQueryloanplansRequest) GetLoanarno() string {
    return r.loanarno
}

func (r *TmallCarLeaseQueryloanplansRequest) SetIproleid(iproleid string) error {
    r.iproleid = iproleid
    r.Set("iproleid", iproleid)
    return nil
}

func (r TmallCarLeaseQueryloanplansRequest) GetIproleid() string {
    return r.iproleid
}

