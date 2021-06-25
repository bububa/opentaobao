package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
优惠券状态更改 APIResponse
alibaba.alsc.crm.voucher.status.change

核销优惠券
*/
type AlibabaAlscCrmVoucherStatusChangeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmVoucherStatusChangeResponse `json:"alibaba_alsc_crm_voucher_status_change_response,omitempty"`
}

type AlibabaAlscCrmVoucherStatusChangeResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
