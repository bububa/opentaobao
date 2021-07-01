package tmic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmic"
)

/* 
获取单题选项 
tmall.tmic.questionnaire.option.get

根据具体题号，获取当前题目的选项列表
*/
func TmallTmicQuestionnaireOptionGet(clt *core.SDKClient, req *tmic.TmallTmicQuestionnaireOptionGetAPIRequest, session string) (*tmic.TmallTmicQuestionnaireOptionGetAPIResponse, error) {
    var resp tmic.TmallTmicQuestionnaireOptionGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
