package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新审批单 APIResponse
alitrip.btrip.approval.update

更新审批单
*/
type AlitripBtripApprovalUpdateAPIResponse struct {
    model.CommonResponse
    AlitripBtripApprovalUpdateResponse
}

type AlitripBtripApprovalUpdateResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_approval_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
