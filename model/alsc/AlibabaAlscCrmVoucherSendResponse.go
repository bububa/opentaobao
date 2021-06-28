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
    AlibabaAlscCrmVoucherSendResponse
}

type AlibabaAlscCrmVoucherSendResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_voucher_send_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
