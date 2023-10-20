package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpWordpackageFindlistAPIRequest 词包列表查询 API请求
// taobao.universalbp.wordpackage.findlist
//
// 根据计划+单元id，查绑定的词包列表
type TaobaoUniversalbpWordpackageFindlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// wordPackageQueryVO
	_wordPackageQueryVO *WordPackageQueryVo
}

// NewTaobaoUniversalbpWordpackageFindlistRequest 初始化TaobaoUniversalbpWordpackageFindlistAPIRequest对象
func NewTaobaoUniversalbpWordpackageFindlistRequest() *TaobaoUniversalbpWordpackageFindlistAPIRequest {
	return &TaobaoUniversalbpWordpackageFindlistAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpWordpackageFindlistAPIRequest) Reset() {
	r._topServiceContext = nil
	r._wordPackageQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpWordpackageFindlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.wordpackage.findlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpWordpackageFindlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpWordpackageFindlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpWordpackageFindlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpWordpackageFindlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetWordPackageQueryVO is WordPackageQueryVO Setter
// wordPackageQueryVO
func (r *TaobaoUniversalbpWordpackageFindlistAPIRequest) SetWordPackageQueryVO(_wordPackageQueryVO *WordPackageQueryVo) error {
	r._wordPackageQueryVO = _wordPackageQueryVO
	r.Set("word_package_query_v_o", _wordPackageQueryVO)
	return nil
}

// GetWordPackageQueryVO WordPackageQueryVO Getter
func (r TaobaoUniversalbpWordpackageFindlistAPIRequest) GetWordPackageQueryVO() *WordPackageQueryVo {
	return r._wordPackageQueryVO
}

var poolTaobaoUniversalbpWordpackageFindlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpWordpackageFindlistRequest()
	},
}

// GetTaobaoUniversalbpWordpackageFindlistRequest 从 sync.Pool 获取 TaobaoUniversalbpWordpackageFindlistAPIRequest
func GetTaobaoUniversalbpWordpackageFindlistAPIRequest() *TaobaoUniversalbpWordpackageFindlistAPIRequest {
	return poolTaobaoUniversalbpWordpackageFindlistAPIRequest.Get().(*TaobaoUniversalbpWordpackageFindlistAPIRequest)
}

// ReleaseTaobaoUniversalbpWordpackageFindlistAPIRequest 将 TaobaoUniversalbpWordpackageFindlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpWordpackageFindlistAPIRequest(v *TaobaoUniversalbpWordpackageFindlistAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpWordpackageFindlistAPIRequest.Put(v)
}
