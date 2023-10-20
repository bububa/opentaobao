package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest 定向推广-新增或删除屏蔽词 API请求
// alibaba.scbp.target.ad.plan.forbidden.word.modify
//
// 定向推广-新增或删除屏蔽词
type AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest struct {
	model.Params
	// TopP4pQuickForbiddenWord
	_topP4pQuickForbiddenWord *TopP4pQuickForbiddenWordDto
}

// NewAlibabaScbpTargetAdPlanForbiddenWordModifyRequest 初始化AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest对象
func NewAlibabaScbpTargetAdPlanForbiddenWordModifyRequest() *AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest {
	return &AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.forbidden.word.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopP4pQuickForbiddenWord is TopP4pQuickForbiddenWord Setter
// TopP4pQuickForbiddenWord
func (r *AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest) SetTopP4pQuickForbiddenWord(_topP4pQuickForbiddenWord *TopP4pQuickForbiddenWordDto) error {
	r._topP4pQuickForbiddenWord = _topP4pQuickForbiddenWord
	r.Set("top_p4p_quick_forbidden_word", _topP4pQuickForbiddenWord)
	return nil
}

// GetTopP4pQuickForbiddenWord TopP4pQuickForbiddenWord Getter
func (r AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest) GetTopP4pQuickForbiddenWord() *TopP4pQuickForbiddenWordDto {
	return r._topP4pQuickForbiddenWord
}
