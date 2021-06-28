package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
储值核销预先校验 APIResponse
alibaba.alsc.crm.recharge.dedutprecheck.get

储值核销预先校验接口
*/
type AlibabaAlscCrmRechargeDedutprecheckGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmRechargeDedutprecheckGetResponse `json:"alibaba_alsc_crm_recharge_dedutprecheck_get_response,omitempty"` 
    AlibabaAlscCrmRechargeDedutprecheckGetResponse
}

/* model for simplify = false
type AlibabaAlscCrmRechargeDedutprecheckGetResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmRechargeDedutprecheckGetResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
