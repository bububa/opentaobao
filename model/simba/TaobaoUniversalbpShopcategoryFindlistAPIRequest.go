package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpShopcategoryFindlistAPIRequest 人群相关类目查询 API请求
// taobao.universalbp.shopcategory.findlist
//
// 查询店铺所属的类目信息
type TaobaoUniversalbpShopcategoryFindlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaoUniversalbpShopcategoryFindlistRequest 初始化TaobaoUniversalbpShopcategoryFindlistAPIRequest对象
func NewTaobaoUniversalbpShopcategoryFindlistRequest() *TaobaoUniversalbpShopcategoryFindlistAPIRequest {
	return &TaobaoUniversalbpShopcategoryFindlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpShopcategoryFindlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.shopcategory.findlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpShopcategoryFindlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpShopcategoryFindlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpShopcategoryFindlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpShopcategoryFindlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}
