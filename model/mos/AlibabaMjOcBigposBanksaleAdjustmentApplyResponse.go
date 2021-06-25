package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
大pos银行卡调账申请 APIResponse
alibaba.mj.oc.bigpos.banksale.adjustment.apply

大pos银行卡调账申请
*/
type AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjOcBigposBanksaleAdjustmentApplyResponse `json:"alibaba_mj_oc_bigpos_banksale_adjustment_apply_response,omitempty"`
}

type AlibabaMjOcBigposBanksaleAdjustmentApplyResponse struct {

}
