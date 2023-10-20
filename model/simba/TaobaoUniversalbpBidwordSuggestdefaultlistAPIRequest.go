package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest 建议默认关键词 API请求
// taobao.universalbp.bidword.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词
type TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// bidwordDefaultQueryVO
	_bidwordDefaultQueryVO *BidwordDefaultQueryVo
}

// NewTaobaoUniversalbpBidwordSuggestdefaultlistRequest 初始化TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest对象
func NewTaobaoUniversalbpBidwordSuggestdefaultlistRequest() *TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest {
	return &TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.bidword.suggestdefaultlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetBidwordDefaultQueryVO is BidwordDefaultQueryVO Setter
// bidwordDefaultQueryVO
func (r *TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest) SetBidwordDefaultQueryVO(_bidwordDefaultQueryVO *BidwordDefaultQueryVo) error {
	r._bidwordDefaultQueryVO = _bidwordDefaultQueryVO
	r.Set("bidword_default_query_v_o", _bidwordDefaultQueryVO)
	return nil
}

// GetBidwordDefaultQueryVO BidwordDefaultQueryVO Getter
func (r TaobaoUniversalbpBidwordSuggestdefaultlistAPIRequest) GetBidwordDefaultQueryVO() *BidwordDefaultQueryVo {
	return r._bidwordDefaultQueryVO
}
