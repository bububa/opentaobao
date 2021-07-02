package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryQueryAPIRequest 后端商品库存查询接口 API请求
// taobao.jst.astrolabe.storeinventory.query
//
// 查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
type TaobaoJstAstrolabeStoreinventoryQueryAPIRequest struct {
	model.Params
	// 门店
	_stores []Store
}

// NewTaobaoJstAstrolabeStoreinventoryQueryRequest 初始化TaobaoJstAstrolabeStoreinventoryQueryAPIRequest对象
func NewTaobaoJstAstrolabeStoreinventoryQueryRequest() *TaobaoJstAstrolabeStoreinventoryQueryAPIRequest {
	return &TaobaoJstAstrolabeStoreinventoryQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStores is Stores Setter
// 门店
func (r *TaobaoJstAstrolabeStoreinventoryQueryAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// GetStores Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryQueryAPIRequest) GetStores() []Store {
	return r._stores
}
