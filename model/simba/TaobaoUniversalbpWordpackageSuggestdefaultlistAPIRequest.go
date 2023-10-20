package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest 建议默认关键词包 API请求
// taobao.universalbp.wordpackage.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词包
type TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// wordPackageDefaultQueryVO
	_wordPackageDefaultQueryVO *WordPackageDefaultQueryVo
}

// NewTaobaouniversalbpwordpackagesuggestdefaultlistRequest 初始化TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest对象
func NewTaobaouniversalbpwordpackagesuggestdefaultlistRequest() *TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest {
	return &TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.wordpackage.suggestdefaultlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetWordPackageDefaultQueryVO is WordPackageDefaultQueryVO Setter
// wordPackageDefaultQueryVO
func (r *TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest) SetWordPackageDefaultQueryVO(_wordPackageDefaultQueryVO *WordPackageDefaultQueryVo) error {
	r._wordPackageDefaultQueryVO = _wordPackageDefaultQueryVO
	r.Set("word_package_default_query_v_o", _wordPackageDefaultQueryVO)
	return nil
}

// GetWordPackageDefaultQueryVO WordPackageDefaultQueryVO Getter
func (r TaobaouniversalbpwordpackagesuggestdefaultlistAPIRequest) GetWordPackageDefaultQueryVO() *WordPackageDefaultQueryVo {
	return r._wordPackageDefaultQueryVO
}
