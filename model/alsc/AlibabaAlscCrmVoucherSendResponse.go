package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发送券给指定用户 APIResponse
alibaba.alsc.crm.voucher.send

发送券给指定用户
*/
type AlibabaAlscCrmVoucherSendAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmVoucherSendResponse `json:"alibaba_alsc_crm_voucher_send_response,omitempty"`
}

type AlibabaAlscCrmVoucherSendResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
