package tmic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmic"
)

// Tmalltmicquestionnaireanswerpush 提交单题答案
// tmall.tmic.questionnaire.answer.push
//
// 问卷单题回答的提交
func Tmalltmicquestionnaireanswerpush(clt *core.SDKClient, req *tmic.TmalltmicquestionnaireanswerpushAPIRequest, session string) (*tmic.TmalltmicquestionnaireanswerpushAPIResponse, error) {
	var resp tmic.TmalltmicquestionnaireanswerpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
