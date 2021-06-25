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
    Response *AlibabaAlscCrmPointChkpntbypayResponse `json:"alibaba_alsc_crm_point_chkpntbypay_response,omitempty"`
}

type AlibabaAlscCrmPointChkpntbypayResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
