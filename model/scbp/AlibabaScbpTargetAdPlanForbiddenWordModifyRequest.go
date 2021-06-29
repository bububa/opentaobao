package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-新增或删除屏蔽词 API请求
alibaba.scbp.target.ad.plan.forbidden.word.modify

定向推广-新增或删除屏蔽词
*/
type AlibabaScbpTargetAdPlanForbiddenWordModifyRequest struct {
    model.Params
    // TopP4pQuickForbiddenWord
    topP4pQuickForbiddenWord   *TopP4pQuickForbiddenWordDto
}

// 初始化AlibabaScbpTargetAdPlanForbiddenWordModifyRequest对象
func NewAlibabaScbpTargetAdPlanForbiddenWordModifyRequest() *AlibabaScbpTargetAdPlanForbiddenWordModifyRequest{
    return &AlibabaScbpTargetAdPlanForbiddenWordModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanForbiddenWordModifyRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.forbidden.word.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanForbiddenWordModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopP4pQuickForbiddenWord Setter
// TopP4pQuickForbiddenWord
func (r *AlibabaScbpTargetAdPlanForbiddenWordModifyRequest) SetTopP4pQuickForbiddenWord(topP4pQuickForbiddenWord *TopP4pQuickForbiddenWordDto) error {
    r.topP4pQuickForbiddenWord = topP4pQuickForbiddenWord
    r.Set("top_p4p_quick_forbidden_word", topP4pQuickForbiddenWord)
    return nil
}

// TopP4pQuickForbiddenWord Getter
func (r AlibabaScbpTargetAdPlanForbiddenWordModifyRequest) GetTopP4pQuickForbiddenWord() *TopP4pQuickForbiddenWordDto {
    return r.topP4pQuickForbiddenWord
}
