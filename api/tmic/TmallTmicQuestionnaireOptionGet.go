package tmic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmic"
)

// Tmalltmicquestionnaireoptionget 获取单题选项
// tmall.tmic.questionnaire.option.get
//
// 根据具体题号，获取当前题目的选项列表
func Tmalltmicquestionnaireoptionget(clt *core.SDKClient, req *tmic.TmalltmicquestionnaireoptiongetAPIRequest, session string) (*tmic.TmalltmicquestionnaireoptiongetAPIResponse, error) {
	var resp tmic.TmalltmicquestionnaireoptiongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
