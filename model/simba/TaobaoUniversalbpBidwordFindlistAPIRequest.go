package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpBidwordFindlistAPIRequest 词列表查询 API请求
// taobao.universalbp.bidword.findlist
//
// 根据计划+单元id，查绑定的词列表
type TaobaoUniversalbpBidwordFindlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// wordQueryVO
	_wordQueryVO *WordQueryVo
}

// NewTaobaoUniversalbpBidwordFindlistRequest 初始化TaobaoUniversalbpBidwordFindlistAPIRequest对象
func NewTaobaoUniversalbpBidwordFindlistRequest() *TaobaoUniversalbpBidwordFindlistAPIRequest {
	return &TaobaoUniversalbpBidwordFindlistAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpBidwordFindlistAPIRequest) Reset() {
	r._topServiceContext = nil
	r._wordQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpBidwordFindlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.bidword.findlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpBidwordFindlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpBidwordFindlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpBidwordFindlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpBidwordFindlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetWordQueryVO is WordQueryVO Setter
// wordQueryVO
func (r *TaobaoUniversalbpBidwordFindlistAPIRequest) SetWordQueryVO(_wordQueryVO *WordQueryVo) error {
	r._wordQueryVO = _wordQueryVO
	r.Set("word_query_v_o", _wordQueryVO)
	return nil
}

// GetWordQueryVO WordQueryVO Getter
func (r TaobaoUniversalbpBidwordFindlistAPIRequest) GetWordQueryVO() *WordQueryVo {
	return r._wordQueryVO
}

var poolTaobaoUniversalbpBidwordFindlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpBidwordFindlistRequest()
	},
}

// GetTaobaoUniversalbpBidwordFindlistRequest 从 sync.Pool 获取 TaobaoUniversalbpBidwordFindlistAPIRequest
func GetTaobaoUniversalbpBidwordFindlistAPIRequest() *TaobaoUniversalbpBidwordFindlistAPIRequest {
	return poolTaobaoUniversalbpBidwordFindlistAPIRequest.Get().(*TaobaoUniversalbpBidwordFindlistAPIRequest)
}

// ReleaseTaobaoUniversalbpBidwordFindlistAPIRequest 将 TaobaoUniversalbpBidwordFindlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpBidwordFindlistAPIRequest(v *TaobaoUniversalbpBidwordFindlistAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpBidwordFindlistAPIRequest.Put(v)
}
