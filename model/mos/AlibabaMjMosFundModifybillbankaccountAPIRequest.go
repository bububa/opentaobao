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
type AlibabaMjMosFundModifybillbankaccountAPIRequest struct {
    model.Params
    // 修改入参
    _modifyDto   *ModifyBillDto
}

// 初始化AlibabaMjMosFundModifybillbankaccountAPIRequest对象
func NewAlibabaMjMosFundModifybillbankaccountRequest() *AlibabaMjMosFundModifybillbankaccountAPIRequest{
    return &AlibabaMjMosFundModifybillbankaccountAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjMosFundModifybillbankaccountAPIRequest) GetApiMethodName() string {
    return "alibaba.mj.mos.fund.modifybillbankaccount"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjMosFundModifybillbankaccountAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ModifyDto Setter
// 修改入参
func (r *AlibabaMjMosFundModifybillbankaccountAPIRequest) SetModifyDto(_modifyDto *ModifyBillDto) error {
    r._modifyDto = _modifyDto
    r.Set("modify_dto", _modifyDto)
    return nil
}

// ModifyDto Getter
func (r AlibabaMjMosFundModifybillbankaccountAPIRequest) GetModifyDto() *ModifyBillDto {
    return r._modifyDto
}
