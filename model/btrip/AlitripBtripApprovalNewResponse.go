package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新建审批单 APIResponse
alitrip.btrip.approval.new

用户新建审批单
*/
type AlitripBtripApprovalNewAPIResponse struct {
    model.CommonResponse
    AlitripBtripApprovalNewResponse
}

type AlitripBtripApprovalNewResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_approval_new_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
