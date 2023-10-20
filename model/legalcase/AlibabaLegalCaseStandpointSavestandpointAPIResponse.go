package legalcase

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcasestandpointsavestandpointAPIResponse 新增反馈口径 API返回值
// alibaba.legal.case.standpoint.savestandpoint
//
// 新增反馈口径 ,从外部接受反馈的口径
type AlibabalegalcasestandpointsavestandpointAPIResponse struct {
	model.CommonResponse
	AlibabalegalcasestandpointsavestandpointAPIResponseModel
}

// AlibabalegalcasestandpointsavestandpointAPIResponseModel is 新增反馈口径 成功返回结果
type AlibabalegalcasestandpointsavestandpointAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_standpoint_savestandpoint_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误编码
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误描述
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 反馈的新增口径id
	Content int64 `json:"content,omitempty" xml:"content,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
