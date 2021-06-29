package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改审批单 API返回值 
alitrip.btrip.approval.modify

修改审批单
*/
type AlitripBtripApprovalModifyAPIResponse struct {
    model.CommonResponse
    AlitripBtripApprovalModifyResponse
}

// 修改审批单 成功返回结果
type AlitripBtripApprovalModifyResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_approval_modify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BtripApplyResult `json:"result,omitempty" xml:"result,omitempty"`
}
