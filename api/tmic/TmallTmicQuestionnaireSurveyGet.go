package tmic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmic"
)

// Tmalltmicquestionnairesurveyget 天猫新品创新中心问卷数据获取
// tmall.tmic.questionnaire.survey.get
//
// 天猫新品创新中心问卷数据获取
func Tmalltmicquestionnairesurveyget(clt *core.SDKClient, req *tmic.TmalltmicquestionnairesurveygetAPIRequest, session string) (*tmic.TmalltmicquestionnairesurveygetAPIResponse, error) {
	var resp tmic.TmalltmicquestionnairesurveygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
