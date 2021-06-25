package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
储值充值 APIResponse
alibaba.alsc.crm.recharge.charge.update

顾客储值账户充值
*/
type AlibabaAlscCrmRechargeChargeUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmRechargeChargeUpdateResponse `json:"alibaba_alsc_crm_recharge_charge_update_response,omitempty"`
}

type AlibabaAlscCrmRechargeChargeUpdateResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
