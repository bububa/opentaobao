package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商银行账号查询 APIResponse
alibaba.mos.finance.bankinfo.querybank

查询供应商对应的银行账号信息
*/
type AlibabaMosFinanceBankinfoQuerybankAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mos_finance_bankinfo_querybank_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *AlibabaMosFinanceBankinfoQuerybankResultDo `json:"result,omitempty" xml:"