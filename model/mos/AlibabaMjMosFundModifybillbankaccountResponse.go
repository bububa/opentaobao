package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改付款单的银行账户信息 APIResponse
alibaba.mj.mos.fund.modifybillbankaccount

修改付款单的银行账户信息
*/
type AlibabaMjMosFundModifybillbankaccountAPIResponse struct {
    model.CommonResponse
    AlibabaMjMosFundModifybillbankaccountResponse
}

type AlibabaMjMosFundModifybillbankaccountResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_mos_fund_modifybillbankaccount_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // data
    
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`

    
}
