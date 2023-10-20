package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbplabelfindconfiglistAPIRequest 查询可用标签id信息 API请求
// taobao.universalbp.label.findconfiglist
//
// 入参账号信息，出参可用标签id，用于下游接口入参
type TaobaouniversalbplabelfindconfiglistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaouniversalbplabelfindconfiglistRequest 初始化TaobaouniversalbplabelfindconfiglistAPIRequest对象
func NewTaobaouniversalbplabelfindconfiglistRequest() *TaobaouniversalbplabelfindconfiglistAPIRequest {
	return &TaobaouniversalbplabelfindconfiglistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbplabelfindconfiglistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.label.findconfiglist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbplabelfindconfiglistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbplabelfindconfiglistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbplabelfindconfiglistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbplabelfindconfiglistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}
