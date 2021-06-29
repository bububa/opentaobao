package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构同步发票信息给阿里健康 APIResponse
alibaba.alihealth.examination.invoice.info.notify

体检机构向阿里健康同步发票信息
*/
type AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationInvoiceInfoNotifyResponse
}

type AlibabaAlihealthExaminationInvoiceInfoNotifyResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_invoice_info_notify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
