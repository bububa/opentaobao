package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应商银行账号查询 APIResponse
alibaba.mos.finance.bankinfo.querybank

查询供应商对应的银行账号信息
*/
type AlibabaMosFinanceBankinfoQuerybankAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMosFinanceBankinfoQuerybankResponse `json:"alibaba_mos_finance_bankinfo_querybank_response,omitempty"`
}

type AlibabaMosFinanceBankinfoQuerybankResponse struct {

    // 返回结果
    Result   *AlibabaMosFinanceBankinfoQuerybankResultDo `json:"result,omitempty"`

}
