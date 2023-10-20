package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest 建议默认关键词 API请求
// taobao.universalbp.bidword.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词
type TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// bidwordDefaultQueryVO
	_bidwordDefaultQueryVO *BidwordDefaultQueryVo
}

// NewTaobaouniversalbpbidwordsuggestdefaultlistRequest 初始化TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest对象
func NewTaobaouniversalbpbidwordsuggestdefaultlistRequest() *TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest {
	return &TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.bidword.suggestdefaultlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetBidwordDefaultQueryVO is BidwordDefaultQueryVO Setter
// bidwordDefaultQueryVO
func (r *TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest) SetBidwordDefaultQueryVO(_bidwordDefaultQueryVO *BidwordDefaultQueryVo) error {
	r._bidwordDefaultQueryVO = _bidwordDefaultQueryVO
	r.Set("bidword_default_query_v_o", _bidwordDefaultQueryVO)
	return nil
}

// GetBidwordDefaultQueryVO BidwordDefaultQueryVO Getter
func (r TaobaouniversalbpbidwordsuggestdefaultlistAPIRequest) GetBidwordDefaultQueryVO() *BidwordDefaultQueryVo {
	return r._bidwordDefaultQueryVO
}
