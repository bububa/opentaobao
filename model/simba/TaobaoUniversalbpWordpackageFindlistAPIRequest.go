package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpwordpackagefindlistAPIRequest 词包列表查询 API请求
// taobao.universalbp.wordpackage.findlist
//
// 根据计划+单元id，查绑定的词包列表
type TaobaouniversalbpwordpackagefindlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// wordPackageQueryVO
	_wordPackageQueryVO *WordPackageQueryVo
}

// NewTaobaouniversalbpwordpackagefindlistRequest 初始化TaobaouniversalbpwordpackagefindlistAPIRequest对象
func NewTaobaouniversalbpwordpackagefindlistRequest() *TaobaouniversalbpwordpackagefindlistAPIRequest {
	return &TaobaouniversalbpwordpackagefindlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpwordpackagefindlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.wordpackage.findlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpwordpackagefindlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpwordpackagefindlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpwordpackagefindlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpwordpackagefindlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetWordPackageQueryVO is WordPackageQueryVO Setter
// wordPackageQueryVO
func (r *TaobaouniversalbpwordpackagefindlistAPIRequest) SetWordPackageQueryVO(_wordPackageQueryVO *WordPackageQueryVo) error {
	r._wordPackageQueryVO = _wordPackageQueryVO
	r.Set("word_package_query_v_o", _wordPackageQueryVO)
	return nil
}

// GetWordPackageQueryVO WordPackageQueryVO Getter
func (r TaobaouniversalbpwordpackagefindlistAPIRequest) GetWordPackageQueryVO() *WordPackageQueryVo {
	return r._wordPackageQueryVO
}
