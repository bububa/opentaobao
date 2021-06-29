package tmic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmic"
)

/* 
天猫新品创新中心问卷数据获取 
tmall.tmic.questionnaire.survey.get

天猫新品创新中心问卷数据获取
*/
func TmallTmicQuestionnaireSurveyGet(clt *core.SDKClient, req *tmic.TmallTmicQuestionnaireSurveyGetRequest, session string) (*tmic.TmallTmicQuestionnaireSurveyGetAPIResponse, error) {
    var resp tmic.TmallTmicQuestionnaireSurveyGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
