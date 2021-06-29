package icburfq

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商提交报价 APIResponse
alibaba.icbu.quotation.post

供应商对RFQ进行报价
*/
type AlibabaIcbuQuotationPostAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuQuotationPostResponse
}

type AlibabaIcbuQuotationPostResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_quotation_post_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求返回结果信息
    
    Result   *RfqRemoteServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
