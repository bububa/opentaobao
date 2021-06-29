package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大pos银行卡调账申请 API返回值 
alibaba.mj.oc.bigpos.banksale.adjustment.apply

大pos银行卡调账申请
*/
type AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse struct {
    model.CommonResponse
    AlibabaMjOcBigposBanksaleAdjustmentApplyResponse
}

// 大pos银行卡调账申请 成功返回结果
type AlibabaMjOcBigposBanksaleAdjustmentApplyResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_oc_bigpos_banksale_adjustment_apply_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
