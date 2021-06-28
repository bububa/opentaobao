package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
校验支付链路中的积分抵扣是否合法 APIResponse
alibaba.alsc.crm.point.chkpntbypay

校验支付链路中的积分抵扣是否合法
*/
type AlibabaAlscCrmPointChkpntbypayAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmPointChkpntbypayResponse `json:"alibaba_alsc_crm_point_chkpntbypay_response,omitempty"` 
    AlibabaAlscCrmPointChkpntbypayResponse
}

/* model for simplify = false
type AlibabaAlscCrmPointChkpntbypayResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmPointChkpntbypayResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
