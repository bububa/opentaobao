package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发送券给指定用户 APIResponse
alibaba.alsc.crm.voucher.send

发送券给指定用户
*/
type AlibabaAlscCrmVoucherSendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_crm_voucher_send_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"