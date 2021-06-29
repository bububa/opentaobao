package icburfq

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询RFQ APIResponse
alibaba.icbu.rfq.search

用于查询RFQ的信息
*/
type AlibabaIcbuRfqSearchAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuRfqSearchResponse
}

type AlibabaIcbuRfqSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_rfq_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回信息结果集
    
    Result   *RfqRemoteServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
