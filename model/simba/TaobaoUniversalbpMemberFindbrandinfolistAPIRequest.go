package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpMemberFindbrandinfolistAPIRequest 查询可用品牌列表 API请求
// taobao.universalbp.member.findbrandinfolist
//
// 查询账号对应的品牌，用于品牌人群屏蔽等
type TaobaoUniversalbpMemberFindbrandinfolistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaoUniversalbpMemberFindbrandinfolistRequest 初始化TaobaoUniversalbpMemberFindbrandinfolistAPIRequest对象
func NewTaobaoUniversalbpMemberFindbrandinfolistRequest() *TaobaoUniversalbpMemberFindbrandinfolistAPIRequest {
	return &TaobaoUniversalbpMemberFindbrandinfolistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpMemberFindbrandinfolistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.member.findbrandinfolist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpMemberFindbrandinfolistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpMemberFindbrandinfolistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpMemberFindbrandinfolistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpMemberFindbrandinfolistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}
