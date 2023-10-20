package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest 关键词包建议 API请求
// taobao.universalbp.wordpackage.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词包
type TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// wordPackageSuggestQueryVO
	_wordPackageSuggestQueryVO *WordPackageSuggestQueryVo
}

// NewTaobaoUniversalbpWordpackageSuggestkrlistRequest 初始化TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest对象
func NewTaobaoUniversalbpWordpackageSuggestkrlistRequest() *TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest {
	return &TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest) Reset() {
	r._topServiceContext = nil
	r._wordPackageSuggestQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.wordpackage.suggestkrlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetWordPackageSuggestQueryVO is WordPackageSuggestQueryVO Setter
// wordPackageSuggestQueryVO
func (r *TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest) SetWordPackageSuggestQueryVO(_wordPackageSuggestQueryVO *WordPackageSuggestQueryVo) error {
	r._wordPackageSuggestQueryVO = _wordPackageSuggestQueryVO
	r.Set("word_package_suggest_query_v_o", _wordPackageSuggestQueryVO)
	return nil
}

// GetWordPackageSuggestQueryVO WordPackageSuggestQueryVO Getter
func (r TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest) GetWordPackageSuggestQueryVO() *WordPackageSuggestQueryVo {
	return r._wordPackageSuggestQueryVO
}

var poolTaobaoUniversalbpWordpackageSuggestkrlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpWordpackageSuggestkrlistRequest()
	},
}

// GetTaobaoUniversalbpWordpackageSuggestkrlistRequest 从 sync.Pool 获取 TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest
func GetTaobaoUniversalbpWordpackageSuggestkrlistAPIRequest() *TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest {
	return poolTaobaoUniversalbpWordpackageSuggestkrlistAPIRequest.Get().(*TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest)
}

// ReleaseTaobaoUniversalbpWordpackageSuggestkrlistAPIRequest 将 TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpWordpackageSuggestkrlistAPIRequest(v *TaobaoUniversalbpWordpackageSuggestkrlistAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpWordpackageSuggestkrlistAPIRequest.Put(v)
}
