package tmic

import (
	"encoding/xml"

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

// TmallTmicQuestionnaireSurveyGetAPIResponseModel is 天猫新品创新中心问卷数据获取 成功返回结果
type TmallTmicQuestionnaireSurveyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_tmic_questionnaire_survey_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallTmicQuestionnaireSurveyGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
