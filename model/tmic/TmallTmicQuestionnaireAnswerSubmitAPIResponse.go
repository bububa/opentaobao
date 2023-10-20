package tmic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTmicQuestionnaireAnswerSubmitAPIResponse 提交问卷答案 API返回值
// tmall.tmic.questionnaire.answer.submit
//
// 天猫新品创新中心对外开放问卷，提交问卷答案
type TmallTmicQuestionnaireAnswerSubmitAPIResponse struct {
	model.CommonResponse
	TmallTmicQuestionnaireAnswerSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTmicQuestionnaireAnswerSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTmicQuestionnaireAnswerSubmitAPIResponseModel).Reset()
}

// TmallTmicQuestionnaireAnswerSubmitAPIResponseModel is 提交问卷答案 成功返回结果
type TmallTmicQuestionnaireAnswerSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_tmic_questionnaire_answer_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示语
	BizErrorInfo string `json:"biz_error_info,omitempty" xml:"biz_error_info,omitempty"`
	// 错误编码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 是否调用成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallTmicQuestionnaireAnswerSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorInfo = ""
	m.BizErrorCode = ""
	m.BizSuccess = false
}

var poolTmallTmicQuestionnaireAnswerSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTmicQuestionnaireAnswerSubmitAPIResponse)
	},
}

// GetTmallTmicQuestionnaireAnswerSubmitAPIResponse 从 sync.Pool 获取 TmallTmicQuestionnaireAnswerSubmitAPIResponse
func GetTmallTmicQuestionnaireAnswerSubmitAPIResponse() *TmallTmicQuestionnaireAnswerSubmitAPIResponse {
	return poolTmallTmicQuestionnaireAnswerSubmitAPIResponse.Get().(*TmallTmicQuestionnaireAnswerSubmitAPIResponse)
}

// ReleaseTmallTmicQuestionnaireAnswerSubmitAPIResponse 将 TmallTmicQuestionnaireAnswerSubmitAPIResponse 保存到 sync.Pool
func ReleaseTmallTmicQuestionnaireAnswerSubmitAPIResponse(v *TmallTmicQuestionnaireAnswerSubmitAPIResponse) {
	v.Reset()
	poolTmallTmicQuestionnaireAnswerSubmitAPIResponse.Put(v)
}
