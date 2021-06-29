package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.issue.partner.rma.screening.create APIResponse
aliexpress.solution.issue.partner.rma.screening.create

Receives information about screening results from after sales partners
*/
type AliexpressSolutionIssuePartnerRmaScreeningCreateAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionIssuePartnerRmaScreeningCreateResponse
}

type AliexpressSolutionIssuePartnerRmaScreeningCreateResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_issue_partner_rma_screening_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // error code. 0 means no error
    
    CodeError   string `json:"code_error,omitempty" xml:"code_error,omitempty"`

    
    // error description
    
    ErrorDescription   string `json:"error_description,omitempty" xml:"error_description,omitempty"`

    
}
