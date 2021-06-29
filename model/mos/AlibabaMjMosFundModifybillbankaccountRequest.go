package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改付款单的银行账户信息 API请求
alibaba.mj.mos.fund.modifybillbankaccount

修改付款单的银行账户信息
*/
type AlibabaMjMosFundModifybillbankaccountRequest struct {
    model.Params
    // 修改入参
    _modifyDto   *ModifyBillDto
}

// 初始化AlibabaMjMosFundModifybillbankaccountRequest对象
func NewAlibabaMjMosFundModifybillbankaccountRequest() *AlibabaMjMosFundModifybillbankaccountRequest{
    return &AlibabaMjMosFundModifybillbankaccountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjMosFundModifybillbankaccountRequest) GetApiMethodName() string {
    return "alibaba.mj.mos.fund.modifybillbankaccount"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjMosFundModifybillbankaccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ModifyDto Setter
// 修改入参
func (r *AlibabaMjMosFundModifybillbankaccountRequest) SetModifyDto(_modifyDto *ModifyBillDto) error {
    r._modifyDto = _modifyDto
    r.Set("modify_dto", _modifyDto)
    return nil
}

// ModifyDto Getter
func (r AlibabaMjMosFundModifybillbankaccountRequest) GetModifyDto() *ModifyBillDto {
    return r._modifyDto
}
