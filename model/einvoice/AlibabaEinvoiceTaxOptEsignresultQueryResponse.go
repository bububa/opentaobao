package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户签约税优结果 APIResponse
alibaba.einvoice.tax.opt.esignresult.query

查询用户是否已经签约
*/
type AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceTaxOptEsignresultQueryResponse
}

type AlibabaEinvoiceTaxOptEsignresultQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_esignresult_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果
    
    Results   []AgreementInfoDto `json:"results,omitempty" xml:"results>agreement_info_dto,omitempty"`
    
    
}
