package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripApprovalUpdateAPIResponse 更新审批单 API返回值
// alitrip.btrip.approval.update
//
// 更新审批单
type AlitripBtripApprovalUpdateAPIResponse struct {
	model.CommonResponse
	AlitripBtripApprovalUpdateAPIResponseModel
}

// AlitripBtripApprovalUpdateAPIResponseModel is 更新审批单 成功返回结果
type AlitripBtripApprovalUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_approval_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}
