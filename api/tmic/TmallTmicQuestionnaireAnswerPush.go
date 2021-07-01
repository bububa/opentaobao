package tmic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmic"
)

/* 
提交单题答案 
tmall.tmic.questionnaire.answer.push

问卷单题回答的提交
*/
func TmallTmicQuestionnaireAnswerPush(clt *core.SDKClient, req *tmic.TmallTmicQuestionnaireAnswerPushAPIRequest, session string) (*tmic.TmallTmicQuestionnaireAnswerPushAPIResponse, error) {
    var resp tmic.TmallTmicQuestionnaireAnswerPushAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
