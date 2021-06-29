package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交发票申请 API返回值 
alibaba.einvoice.prod.apply

提交开票申请，如果商户授权自动开票则自动转开票，否则等待商户审核。
*/
type AlibabaEinvoiceProdApplyAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceProdApplyResponse
}

// 提交发票申请 成功返回结果
type AlibabaEinvoiceProdApplyResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_prod_apply_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
