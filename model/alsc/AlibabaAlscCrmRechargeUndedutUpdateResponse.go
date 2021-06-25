package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
储值消费退款(逆向) APIResponse
alibaba.alsc.crm.recharge.undedut.update

新增储值消费退款接口
*/
type AlibabaAlscCrmRechargeUndedutUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmRechargeUndedutUpdateResponse `json:"alibaba_alsc_crm_recharge_undedut_update_response,omitempty"`
}

type AlibabaAlscCrmRechargeUndedutUpdateResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
