package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
储值消费 APIResponse
alibaba.alsc.crm.recharge.dedut.update

增加储值消费接口
*/
type AlibabaAlscCrmRechargeDedutUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmRechargeDedutUpdateResponse `json:"alibaba_alsc_crm_recharge_dedut_update_response,omitempty"`
}

type AlibabaAlscCrmRechargeDedutUpdateResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
