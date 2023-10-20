package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest 建议默认关键词包 API请求
// taobao.universalbp.wordpackage.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词包
type TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// wordPackageDefaultQueryVO
	_wordPackageDefaultQueryVO *WordPackageDefaultQueryVo
}

// NewTaobaoUniversalbpWordpackageSuggestdefaultlistRequest 初始化TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest对象
func NewTaobaoUniversalbpWordpackageSuggestdefaultlistRequest() *TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest {
	return &TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest) Reset() {
	r._topServiceContext = nil
	r._wordPackageDefaultQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.wordpackage.suggestdefaultlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetWordPackageDefaultQueryVO is WordPackageDefaultQueryVO Setter
// wordPackageDefaultQueryVO
func (r *TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest) SetWordPackageDefaultQueryVO(_wordPackageDefaultQueryVO *WordPackageDefaultQueryVo) error {
	r._wordPackageDefaultQueryVO = _wordPackageDefaultQueryVO
	r.Set("word_package_default_query_v_o", _wordPackageDefaultQueryVO)
	return nil
}

// GetWordPackageDefaultQueryVO WordPackageDefaultQueryVO Getter
func (r TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest) GetWordPackageDefaultQueryVO() *WordPackageDefaultQueryVo {
	return r._wordPackageDefaultQueryVO
}

var poolTaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpWordpackageSuggestdefaultlistRequest()
	},
}

// GetTaobaoUniversalbpWordpackageSuggestdefaultlistRequest 从 sync.Pool 获取 TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest
func GetTaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest() *TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest {
	return poolTaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest.Get().(*TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest)
}

// ReleaseTaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest 将 TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest(v *TaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpWordpackageSuggestdefaultlistAPIRequest.Put(v)
}
