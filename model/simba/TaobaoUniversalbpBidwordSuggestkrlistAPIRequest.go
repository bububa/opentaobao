package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpbidwordsuggestkrlistAPIRequest 关键词建议 API请求
// taobao.universalbp.bidword.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词
type TaobaouniversalbpbidwordsuggestkrlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// bidwordSuggestQueryVO
	_bidwordSuggestQueryVO *BidwordSuggestQueryVo
}

// NewTaobaouniversalbpbidwordsuggestkrlistRequest 初始化TaobaouniversalbpbidwordsuggestkrlistAPIRequest对象
func NewTaobaouniversalbpbidwordsuggestkrlistRequest() *TaobaouniversalbpbidwordsuggestkrlistAPIRequest {
	return &TaobaouniversalbpbidwordsuggestkrlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpbidwordsuggestkrlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.bidword.suggestkrlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpbidwordsuggestkrlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpbidwordsuggestkrlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpbidwordsuggestkrlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpbidwordsuggestkrlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetBidwordSuggestQueryVO is BidwordSuggestQueryVO Setter
// bidwordSuggestQueryVO
func (r *TaobaouniversalbpbidwordsuggestkrlistAPIRequest) SetBidwordSuggestQueryVO(_bidwordSuggestQueryVO *BidwordSuggestQueryVo) error {
	r._bidwordSuggestQueryVO = _bidwordSuggestQueryVO
	r.Set("bidword_suggest_query_v_o", _bidwordSuggestQueryVO)
	return nil
}

// GetBidwordSuggestQueryVO BidwordSuggestQueryVO Getter
func (r TaobaouniversalbpbidwordsuggestkrlistAPIRequest) GetBidwordSuggestQueryVO() *BidwordSuggestQueryVo {
	return r._bidwordSuggestQueryVO
}
