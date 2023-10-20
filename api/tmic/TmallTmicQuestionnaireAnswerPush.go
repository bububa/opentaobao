package tmic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmic"
)

// TmallTmicQuestionnaireAnswerPush 提交单题答案
// tmall.tmic.questionnaire.answer.push
//
// 问卷单题回答的提交
func TmallTmicQuestionnaireAnswerPush(clt *core.SDKClient, req *tmic.TmallTmicQuestionnaireAnswerPushAPIRequest, resp *tmic.TmallTmicQuestionnaireAnswerPushAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
