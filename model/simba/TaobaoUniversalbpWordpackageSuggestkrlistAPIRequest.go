package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpwordpackagesuggestkrlistAPIRequest 关键词包建议 API请求
// taobao.universalbp.wordpackage.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词包
type TaobaouniversalbpwordpackagesuggestkrlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// wordPackageSuggestQueryVO
	_wordPackageSuggestQueryVO *WordPackageSuggestQueryVo
}

// NewTaobaouniversalbpwordpackagesuggestkrlistRequest 初始化TaobaouniversalbpwordpackagesuggestkrlistAPIRequest对象
func NewTaobaouniversalbpwordpackagesuggestkrlistRequest() *TaobaouniversalbpwordpackagesuggestkrlistAPIRequest {
	return &TaobaouniversalbpwordpackagesuggestkrlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpwordpackagesuggestkrlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.wordpackage.suggestkrlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpwordpackagesuggestkrlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpwordpackagesuggestkrlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpwordpackagesuggestkrlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpwordpackagesuggestkrlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetWordPackageSuggestQueryVO is WordPackageSuggestQueryVO Setter
// wordPackageSuggestQueryVO
func (r *TaobaouniversalbpwordpackagesuggestkrlistAPIRequest) SetWordPackageSuggestQueryVO(_wordPackageSuggestQueryVO *WordPackageSuggestQueryVo) error {
	r._wordPackageSuggestQueryVO = _wordPackageSuggestQueryVO
	r.Set("word_package_suggest_query_v_o", _wordPackageSuggestQueryVO)
	return nil
}

// GetWordPackageSuggestQueryVO WordPackageSuggestQueryVO Getter
func (r TaobaouniversalbpwordpackagesuggestkrlistAPIRequest) GetWordPackageSuggestQueryVO() *WordPackageSuggestQueryVo {
	return r._wordPackageSuggestQueryVO
}
