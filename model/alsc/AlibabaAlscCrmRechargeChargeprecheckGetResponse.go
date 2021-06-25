package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
储值账户充值前校验 APIResponse
alibaba.alsc.crm.recharge.chargeprecheck.get

储值账户充值前校验接口
*/
type AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmRechargeChargeprecheckGetResponse `json:"alibaba_alsc_crm_recharge_chargeprecheck_get_response,omitempty"`
}

type AlibabaAlscCrmRechargeChargeprecheckGetResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
