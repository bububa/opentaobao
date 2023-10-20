package legalcase

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcasestandpointfeedbackAPIResponse 新增或更新 反馈口径(采纳口径/不采纳口径) API返回值
// alibaba.legal.case.standpoint.feedback
//
// 新增或更新 反馈口径(采纳口径/不采纳口径)
type AlibabalegalcasestandpointfeedbackAPIResponse struct {
	model.CommonResponse
	AlibabalegalcasestandpointfeedbackAPIResponseModel
}

// AlibabalegalcasestandpointfeedbackAPIResponseModel is 新增或更新 反馈口径(采纳口径/不采纳口径) 成功返回结果
type AlibabalegalcasestandpointfeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_standpoint_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误编码
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误信息描述
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 成功失败 true/false
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}
