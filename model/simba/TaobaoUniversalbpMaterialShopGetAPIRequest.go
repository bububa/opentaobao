package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpMaterialShopGetAPIRequest 获取店铺信息 API请求
// taobao.universalbp.material.shop.get
//
// 获取店铺信息
type TaobaoUniversalbpMaterialShopGetAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaoUniversalbpMaterialShopGetRequest 初始化TaobaoUniversalbpMaterialShopGetAPIRequest对象
func NewTaobaoUniversalbpMaterialShopGetRequest() *TaobaoUniversalbpMaterialShopGetAPIRequest {
	return &TaobaoUniversalbpMaterialShopGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpMaterialShopGetAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.material.shop.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpMaterialShopGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpMaterialShopGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpMaterialShopGetAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpMaterialShopGetAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}
