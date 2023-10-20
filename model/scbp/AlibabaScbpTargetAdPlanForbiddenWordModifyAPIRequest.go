package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadplanforbiddenwordmodifyAPIRequest 定向推广-新增或删除屏蔽词 API请求
// alibaba.scbp.target.ad.plan.forbidden.word.modify
//
// 定向推广-新增或删除屏蔽词
type AlibabascbptargetadplanforbiddenwordmodifyAPIRequest struct {
	model.Params
	// TopP4pQuickForbiddenWord
	_topP4pQuickForbiddenWord *TopP4pQuickForbiddenWordDto
}

// NewAlibabascbptargetadplanforbiddenwordmodifyRequest 初始化AlibabascbptargetadplanforbiddenwordmodifyAPIRequest对象
func NewAlibabascbptargetadplanforbiddenwordmodifyRequest() *AlibabascbptargetadplanforbiddenwordmodifyAPIRequest {
	return &AlibabascbptargetadplanforbiddenwordmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptargetadplanforbiddenwordmodifyAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.forbidden.word.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptargetadplanforbiddenwordmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptargetadplanforbiddenwordmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopP4pQuickForbiddenWord is TopP4pQuickForbiddenWord Setter
// TopP4pQuickForbiddenWord
func (r *AlibabascbptargetadplanforbiddenwordmodifyAPIRequest) SetTopP4pQuickForbiddenWord(_topP4pQuickForbiddenWord *TopP4pQuickForbiddenWordDto) error {
	r._topP4pQuickForbiddenWord = _topP4pQuickForbiddenWord
	r.Set("top_p4p_quick_forbidden_word", _topP4pQuickForbiddenWord)
	return nil
}

// GetTopP4pQuickForbiddenWord TopP4pQuickForbiddenWord Getter
func (r AlibabascbptargetadplanforbiddenwordmodifyAPIRequest) GetTopP4pQuickForbiddenWord() *TopP4pQuickForbiddenWordDto {
	return r._topP4pQuickForbiddenWord
}
