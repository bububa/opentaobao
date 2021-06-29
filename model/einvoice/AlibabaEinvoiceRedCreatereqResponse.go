package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发票冲红接口 APIResponse
alibaba.einvoice.red.createreq

发票冲红接口，通过蓝票流水号或者发票号码+发票代码进行冲红
*/
type AlibabaEinvoiceRedCreatereqAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceRedCreatereqResponse
}

type AlibabaEinvoiceRedCreatereqResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_red_createreq_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否冲红成功
    
    IsSuccess   string `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
