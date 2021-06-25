package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-新增或删除屏蔽词 APIRequest
alibaba.scbp.target.ad.plan.forbidden.word.modify

定向推广-新增或删除屏蔽词
*/
type AlibabaScbpTargetAdPlanForbiddenWordModifyRequest struct {
    model.Params

    // TopP4pQuickForbiddenWord
    topP4pQuickForbiddenWord   *TopP4pQuickForbiddenWordDto 

}

func NewAlibabaScbpTargetAdPlanForbiddenWordModifyRequest() *AlibabaScbpTargetAdPlanForbiddenWordModifyRequest{
    return &AlibabaScbpTargetAdPlanForbiddenWordModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdPlanForbiddenWordModifyRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.forbidden.word.modify"
}

func (r AlibabaScbpTargetAdPlanForbiddenWordModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTargetAdPlanForbiddenWordModifyRequest) SetTopP4pQuickForbiddenWord(topP4pQuickForbiddenWord *TopP4pQuickForbiddenWordDto) error {
    r.topP4pQuickForbiddenWord = topP4pQuickForbiddenWord
    r.Set("top_p4p_quick_forbidden_word", topP4pQuickForbiddenWord)
    return nil
}

func (r AlibabaScbpTargetAdPlanForbiddenWordModifyRequest) GetTopP4pQuickForbiddenWord() *TopP4pQuickForbiddenWordDto {
    return r.topP4pQuickForbiddenWord
}

