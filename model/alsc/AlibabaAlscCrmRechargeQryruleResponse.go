package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
储值规则下行 APIResponse
alibaba.alsc.crm.recharge.qryrule

储值规则下行
*/
type AlibabaAlscCrmRechargeQryruleAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmRechargeQryruleResponse `json:"alibaba_alsc_crm_recharge_qryrule_response,omitempty"`
}

type AlibabaAlscCrmRechargeQryruleResponse struct {

    // 返回模型
    Result   *CommonResult `json:"result,omitempty"`

}
