package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCreativePreaddAPIRequest 创建单品创意前置信息 API请求
// taobao.universalbp.creative.preadd
//
// 用于关键词场景创建单品创意前使用
type TaobaoUniversalbpCreativePreaddAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// long
	_tpLong int64
}

// NewTaobaoUniversalbpCreativePreaddRequest 初始化TaobaoUniversalbpCreativePreaddAPIRequest对象
func NewTaobaoUniversalbpCreativePreaddRequest() *TaobaoUniversalbpCreativePreaddAPIRequest {
	return &TaobaoUniversalbpCreativePreaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpCreativePreaddAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.creative.preadd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpCreativePreaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpCreativePreaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpCreativePreaddAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpCreativePreaddAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTpLong is TpLong Setter
// long
func (r *TaobaoUniversalbpCreativePreaddAPIRequest) SetTpLong(_tpLong int64) error {
	r._tpLong = _tpLong
	r.Set("tp_long", _tpLong)
	return nil
}

// GetTpLong TpLong Getter
func (r TaobaoUniversalbpCreativePreaddAPIRequest) GetTpLong() int64 {
	return r._tpLong
}
