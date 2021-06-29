package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新建审批单 API返回值 
alitrip.btrip.approval.new

用户新建审批单
*/
type AlitripBtripApprovalNewAPIResponse struct {
    model.CommonResponse
    AlitripBtripApprovalNewResponse
}

// 新建审批单 成功返回结果
type AlitripBtripApprovalNewResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_approval_new_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}
