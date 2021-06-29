package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.issue.partner.rma.reverselogistic.state.update APIResponse
aliexpress.solution.issue.partner.rma.reverselogistic.state.update

Updates the reverse logistics state for after sales services
*/
type AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateResponse
}

type AliexpressSolutionIssuePartnerRmaReverselogisticStateUpdateResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_issue_partner_rma_reverselogistic_state_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // Error code. 0 value is no error.
    
    CodeError   string `json:"code_error,omitempty" xml:"code_error,omitempty"`

    
    // Error description
    
    ErrorDescription   string `json:"error_description,omitempty" xml:"error_description,omitempty"`

    
}
