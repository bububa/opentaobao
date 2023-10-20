package tmic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmic"
)

// TmallTmicQuestionnaireSurveyGet 天猫新品创新中心问卷数据获取
// tmall.tmic.questionnaire.survey.get
//
// 天猫新品创新中心问卷数据获取
func TmallTmicQuestionnaireSurveyGet(clt *core.SDKClient, req *tmic.TmallTmicQuestionnaireSurveyGetAPIRequest, resp *tmic.TmallTmicQuestionnaireSurveyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
