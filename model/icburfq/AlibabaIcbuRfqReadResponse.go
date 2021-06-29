package icburfq

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
是否已读RFQ APIResponse
alibaba.icbu.rfq.read

是否已读RFQ
*/
type AlibabaIcbuRfqReadAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuRfqReadResponse
}

type AlibabaIcbuRfqReadResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_rfq_read_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
