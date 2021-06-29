package tmic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmic"
)

/* 
提交问卷答案 
tmall.tmic.questionnaire.answer.submit

天猫新品创新中心对外开放问卷，提交问卷答案
*/
func TmallTmicQuestionnaireAnswerSubmit(clt *core.SDKClient, req *tmic.TmallTmicQuestionnaireAnswerSubmitRequest, session string) (*tmic.TmallTmicQuestionnaireAnswerSubmitAPIResponse, error) {
    var resp tmic.TmallTmicQuestionnaireAnswerSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
