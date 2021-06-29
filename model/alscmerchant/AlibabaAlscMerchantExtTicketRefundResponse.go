package alscmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
口碑凭证售后退 APIResponse
alibaba.alsc.merchant.ext.ticket.refund

口碑凭证售后退
*/
type AlibabaAlscMerchantExtTicketRefundAPIResponse struct {
    model.CommonResponse
    AlibabaAlscMerchantExtTicketRefundResponse
}

type AlibabaAlscMerchantExtTicketRefundResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_merchant_ext_ticket_refund_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 外部请求号，支持英文字母和数字，由开发者自行定义（不允许重复）
    
    TicketRequestId   string `json:"ticket_request_id,omitempty" xml:"ticket_request_id,omitempty"`

    
    // 业务返回code
    
    BizCode   string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`

    
}
