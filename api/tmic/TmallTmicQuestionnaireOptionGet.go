package tmic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmic"
)

// TmallTmicQuestionnaireOptionGet 获取单题选项
// tmall.tmic.questionnaire.option.get
//
// 根据具体题号，获取当前题目的选项列表
func TmallTmicQuestionnaireOptionGet(ctx context.Context, clt *core.SDKClient, req *tmic.TmallTmicQuestionnaireOptionGetAPIRequest, resp *tmic.TmallTmicQuestionnaireOptionGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
