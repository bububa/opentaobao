package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpBidwordSuggestkrlistAPIRequest 关键词建议 API请求
// taobao.universalbp.bidword.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词
type TaobaoUniversalbpBidwordSuggestkrlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// bidwordSuggestQueryVO
	_bidwordSuggestQueryVO *BidwordSuggestQueryVo
}

// NewTaobaoUniversalbpBidwordSuggestkrlistRequest 初始化TaobaoUniversalbpBidwordSuggestkrlistAPIRequest对象
func NewTaobaoUniversalbpBidwordSuggestkrlistRequest() *TaobaoUniversalbpBidwordSuggestkrlistAPIRequest {
	return &TaobaoUniversalbpBidwordSuggestkrlistAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpBidwordSuggestkrlistAPIRequest) Reset() {
	r._topServiceContext = nil
	r._bidwordSuggestQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpBidwordSuggestkrlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.bidword.suggestkrlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpBidwordSuggestkrlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpBidwordSuggestkrlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpBidwordSuggestkrlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpBidwordSuggestkrlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetBidwordSuggestQueryVO is BidwordSuggestQueryVO Setter
// bidwordSuggestQueryVO
func (r *TaobaoUniversalbpBidwordSuggestkrlistAPIRequest) SetBidwordSuggestQueryVO(_bidwordSuggestQueryVO *BidwordSuggestQueryVo) error {
	r._bidwordSuggestQueryVO = _bidwordSuggestQueryVO
	r.Set("bidword_suggest_query_v_o", _bidwordSuggestQueryVO)
	return nil
}

// GetBidwordSuggestQueryVO BidwordSuggestQueryVO Getter
func (r TaobaoUniversalbpBidwordSuggestkrlistAPIRequest) GetBidwordSuggestQueryVO() *BidwordSuggestQueryVo {
	return r._bidwordSuggestQueryVO
}

var poolTaobaoUniversalbpBidwordSuggestkrlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpBidwordSuggestkrlistRequest()
	},
}

// GetTaobaoUniversalbpBidwordSuggestkrlistRequest 从 sync.Pool 获取 TaobaoUniversalbpBidwordSuggestkrlistAPIRequest
func GetTaobaoUniversalbpBidwordSuggestkrlistAPIRequest() *TaobaoUniversalbpBidwordSuggestkrlistAPIRequest {
	return poolTaobaoUniversalbpBidwordSuggestkrlistAPIRequest.Get().(*TaobaoUniversalbpBidwordSuggestkrlistAPIRequest)
}

// ReleaseTaobaoUniversalbpBidwordSuggestkrlistAPIRequest 将 TaobaoUniversalbpBidwordSuggestkrlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpBidwordSuggestkrlistAPIRequest(v *TaobaoUniversalbpBidwordSuggestkrlistAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpBidwordSuggestkrlistAPIRequest.Put(v)
}
