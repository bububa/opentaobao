package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpshopcategoryfindlistAPIRequest 人群相关类目查询 API请求
// taobao.universalbp.shopcategory.findlist
//
// 查询店铺所属的类目信息
type TaobaouniversalbpshopcategoryfindlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaouniversalbpshopcategoryfindlistRequest 初始化TaobaouniversalbpshopcategoryfindlistAPIRequest对象
func NewTaobaouniversalbpshopcategoryfindlistRequest() *TaobaouniversalbpshopcategoryfindlistAPIRequest {
	return &TaobaouniversalbpshopcategoryfindlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpshopcategoryfindlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.shopcategory.findlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpshopcategoryfindlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpshopcategoryfindlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpshopcategoryfindlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpshopcategoryfindlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}
