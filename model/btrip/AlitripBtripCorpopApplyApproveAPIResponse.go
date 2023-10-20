package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplyApproveAPIResponse 【商旅】更新审批单状态 API返回值
// alitrip.btrip.corpop.apply.approve
//
// 【商旅】更新审批单状态
type AlitripBtripCorpopApplyApproveAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopApplyApproveAPIResponseModel
}

// AlitripBtripCorpopApplyApproveAPIResponseModel is 【商旅】更新审批单状态 成功返回结果
type AlitripBtripCorpopApplyApproveAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_approve_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 传参
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// 成功
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 0
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 成功标识
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
