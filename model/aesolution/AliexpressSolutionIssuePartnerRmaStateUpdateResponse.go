package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.issue.partner.rma.state.update APIResponse
aliexpress.solution.issue.partner.rma.state.update

Receive changes in state updates for RMAs orders from after sales partners
*/
type AliexpressSolutionIssuePartnerRmaStateUpdateAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionIssuePartnerRmaStateUpdateResponse
}

type AliexpressSolutionIssuePartnerRmaStateUpdateResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_issue_partner_rma_state_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // error code. 0 means no error
    
    CodeError   string `json:"code_error,omitempty" xml:"code_error,omitempty"`

    
    // error description
    
    ErrorDescription   string `json:"error_description,omitempty" xml:"error_description,omitempty"`

    
}
