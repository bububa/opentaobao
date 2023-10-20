package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpmemberfindbrandinfolistAPIRequest 查询可用品牌列表 API请求
// taobao.universalbp.member.findbrandinfolist
//
// 查询账号对应的品牌，用于品牌人群屏蔽等
type TaobaouniversalbpmemberfindbrandinfolistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaouniversalbpmemberfindbrandinfolistRequest 初始化TaobaouniversalbpmemberfindbrandinfolistAPIRequest对象
func NewTaobaouniversalbpmemberfindbrandinfolistRequest() *TaobaouniversalbpmemberfindbrandinfolistAPIRequest {
	return &TaobaouniversalbpmemberfindbrandinfolistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpmemberfindbrandinfolistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.member.findbrandinfolist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpmemberfindbrandinfolistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpmemberfindbrandinfolistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpmemberfindbrandinfolistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpmemberfindbrandinfolistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}
