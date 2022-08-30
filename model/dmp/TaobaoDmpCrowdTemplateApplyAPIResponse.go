package dmp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDmpCrowdTemplateApplyAPIResponse 人群模版采纳并生成人群API API返回值
// taobao.dmp.crowd.template.apply
//
// 人群模版采纳并生成人群API
type TaobaoDmpCrowdTemplateApplyAPIResponse struct {
	model.CommonResponse
	TaobaoDmpCrowdTemplateApplyAPIResponseModel
}

// TaobaoDmpCrowdTemplateApplyAPIResponseModel is 人群模版采纳并生成人群API 成功返回结果
type TaobaoDmpCrowdTemplateApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"dmp_crowd_template_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultErrorCode string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 生成的人群id
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
