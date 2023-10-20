package tmic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTmicQuestionnaireOptionGetAPIResponse 获取单题选项 API返回值
// tmall.tmic.questionnaire.option.get
//
// 根据具体题号，获取当前题目的选项列表
type TmallTmicQuestionnaireOptionGetAPIResponse struct {
	model.CommonResponse
	TmallTmicQuestionnaireOptionGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTmicQuestionnaireOptionGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTmicQuestionnaireOptionGetAPIResponseModel).Reset()
}

// TmallTmicQuestionnaireOptionGetAPIResponseModel is 获取单题选项 成功返回结果
type TmallTmicQuestionnaireOptionGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_tmic_questionnaire_option_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OpenOptionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTmicQuestionnaireOptionGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTmicQuestionnaireOptionGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTmicQuestionnaireOptionGetAPIResponse)
	},
}

// GetTmallTmicQuestionnaireOptionGetAPIResponse 从 sync.Pool 获取 TmallTmicQuestionnaireOptionGetAPIResponse
func GetTmallTmicQuestionnaireOptionGetAPIResponse() *TmallTmicQuestionnaireOptionGetAPIResponse {
	return poolTmallTmicQuestionnaireOptionGetAPIResponse.Get().(*TmallTmicQuestionnaireOptionGetAPIResponse)
}

// ReleaseTmallTmicQuestionnaireOptionGetAPIResponse 将 TmallTmicQuestionnaireOptionGetAPIResponse 保存到 sync.Pool
func ReleaseTmallTmicQuestionnaireOptionGetAPIResponse(v *TmallTmicQuestionnaireOptionGetAPIResponse) {
	v.Reset()
	poolTmallTmicQuestionnaireOptionGetAPIResponse.Put(v)
}
