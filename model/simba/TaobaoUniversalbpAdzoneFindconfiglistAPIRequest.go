package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpadzonefindconfiglistAPIRequest 查询所有可用资源包信息 API请求
// taobao.universalbp.adzone.findconfiglist
//
// 查询该场景下，所有可用的资源包，可能存在数据重复。但是针对不同子场景和推广设置，可以选用的资源包有差异，建议关注补充文档，或者根据bp前端的限制，进行传参。
type TaobaouniversalbpadzonefindconfiglistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaouniversalbpadzonefindconfiglistRequest 初始化TaobaouniversalbpadzonefindconfiglistAPIRequest对象
func NewTaobaouniversalbpadzonefindconfiglistRequest() *TaobaouniversalbpadzonefindconfiglistAPIRequest {
	return &TaobaouniversalbpadzonefindconfiglistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpadzonefindconfiglistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.adzone.findconfiglist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpadzonefindconfiglistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpadzonefindconfiglistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpadzonefindconfiglistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpadzonefindconfiglistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}
