package tmic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmic"
)

// TmallTmicQuestionnaireAnswerSubmit 提交问卷答案
// tmall.tmic.questionnaire.answer.submit
//
// 天猫新品创新中心对外开放问卷，提交问卷答案
func TmallTmicQuestionnaireAnswerSubmit(ctx context.Context, clt *core.SDKClient, req *tmic.TmallTmicQuestionnaireAnswerSubmitAPIRequest, resp *tmic.TmallTmicQuestionnaireAnswerSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
