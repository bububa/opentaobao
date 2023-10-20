package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIRequest 查询dmp浮层配置 API请求
// taobao.universalbp.label.dmp.finddmpmoduleconfig
//
// 入参账号信息，出参达摩盘相关配置信息
type TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaouniversalbplabeldmpfinddmpmoduleconfigRequest 初始化TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIRequest对象
func NewTaobaouniversalbplabeldmpfinddmpmoduleconfigRequest() *TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIRequest {
	return &TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.label.dmp.finddmpmoduleconfig"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}
