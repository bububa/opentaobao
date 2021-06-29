package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退订工单(入驻、加盘、续约) API返回值 
alibaba.einvoice.flow.refund

电子发票工单系统，工单退订能力开放
*/
type AlibabaEinvoiceFlowRefundAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceFlowRefundResponse
}

// 退订工单(入驻、加盘、续约) 成功返回结果
type AlibabaEinvoiceFlowRefundResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_flow_refund_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回参数
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
