package tmic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmic"
)

// Tmalltmicquestionnaireanswersubmit 提交问卷答案
// tmall.tmic.questionnaire.answer.submit
//
// 天猫新品创新中心对外开放问卷，提交问卷答案
func Tmalltmicquestionnaireanswersubmit(clt *core.SDKClient, req *tmic.TmalltmicquestionnaireanswersubmitAPIRequest, session string) (*tmic.TmalltmicquestionnaireanswersubmitAPIResponse, error) {
	var resp tmic.TmalltmicquestionnaireanswersubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
