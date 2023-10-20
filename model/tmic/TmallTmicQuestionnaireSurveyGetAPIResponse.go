package tmic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTmicQuestionnaireSurveyGetAPIResponse 天猫新品创新中心问卷数据获取 API返回值
// tmall.tmic.questionnaire.survey.get
//
// 天猫新品创新中心问卷数据获取
type TmallTmicQuestionnaireSurveyGetAPIResponse struct {
	model.CommonResponse
	TmallTmicQuestionnaireSurveyGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTmicQuestionnaireSurveyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTmicQuestionnaireSurveyGetAPIResponseModel).Reset()
}

// TmallTmicQuestionnaireSurveyGetAPIResponseModel is 天猫新品创新中心问卷数据获取 成功返回结果
type TmallTmicQuestionnaireSurveyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_tmic_questionnaire_survey_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallTmicQuestionnaireSurveyGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTmicQuestionnaireSurveyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTmicQuestionnaireSurveyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTmicQuestionnaireSurveyGetAPIResponse)
	},
}

// GetTmallTmicQuestionnaireSurveyGetAPIResponse 从 sync.Pool 获取 TmallTmicQuestionnaireSurveyGetAPIResponse
func GetTmallTmicQuestionnaireSurveyGetAPIResponse() *TmallTmicQuestionnaireSurveyGetAPIResponse {
	return poolTmallTmicQuestionnaireSurveyGetAPIResponse.Get().(*TmallTmicQuestionnaireSurveyGetAPIResponse)
}

// ReleaseTmallTmicQuestionnaireSurveyGetAPIResponse 将 TmallTmicQuestionnaireSurveyGetAPIResponse 保存到 sync.Pool
func ReleaseTmallTmicQuestionnaireSurveyGetAPIResponse(v *TmallTmicQuestionnaireSurveyGetAPIResponse) {
	v.Reset()
	poolTmallTmicQuestionnaireSurveyGetAPIResponse.Put(v)
}
