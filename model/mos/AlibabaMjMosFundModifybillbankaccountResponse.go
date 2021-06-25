package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改付款单的银行账户信息 APIResponse
alibaba.mj.mos.fund.modifybillbankaccount

修改付款单的银行账户信息
*/
type AlibabaMjMosFundModifybillbankaccountAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjMosFundModifybillbankaccountResponse `json:"alibaba_mj_mos_fund_modifybillbankaccount_response,omitempty"`
}

type AlibabaMjMosFundModifybillbankaccountResponse struct {

    // data
    Data   bool `json:"data,omitempty"`

}
