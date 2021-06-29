package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大pos银行卡调账申请 APIResponse
alibaba.mj.oc.bigpos.banksale.adjustment.apply

大pos银行卡调账申请
*/
type AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse struct {
    model.CommonResponse
    AlibabaMjOcBigposBanksaleAdjustmentApplyResponse
}

type AlibabaMjOcBigposBanksaleAdjustmentApplyResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_oc_bigpos_banksale_adjustment_apply_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

}
