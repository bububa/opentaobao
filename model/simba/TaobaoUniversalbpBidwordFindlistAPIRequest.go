package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpbidwordfindlistAPIRequest 词列表查询 API请求
// taobao.universalbp.bidword.findlist
//
// 根据计划+单元id，查绑定的词列表
type TaobaouniversalbpbidwordfindlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// wordQueryVO
	_wordQueryVO *WordQueryVo
}

// NewTaobaouniversalbpbidwordfindlistRequest 初始化TaobaouniversalbpbidwordfindlistAPIRequest对象
func NewTaobaouniversalbpbidwordfindlistRequest() *TaobaouniversalbpbidwordfindlistAPIRequest {
	return &TaobaouniversalbpbidwordfindlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpbidwordfindlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.bidword.findlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpbidwordfindlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpbidwordfindlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpbidwordfindlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpbidwordfindlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetWordQueryVO is WordQueryVO Setter
// wordQueryVO
func (r *TaobaouniversalbpbidwordfindlistAPIRequest) SetWordQueryVO(_wordQueryVO *WordQueryVo) error {
	r._wordQueryVO = _wordQueryVO
	r.Set("word_query_v_o", _wordQueryVO)
	return nil
}

// GetWordQueryVO WordQueryVO Getter
func (r TaobaouniversalbpbidwordfindlistAPIRequest) GetWordQueryVO() *WordQueryVo {
	return r._wordQueryVO
}
