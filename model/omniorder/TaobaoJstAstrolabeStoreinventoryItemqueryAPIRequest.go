package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest 库存查询接口 API请求
// taobao.jst.astrolabe.storeinventory.itemquery
//
// 查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
type TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest struct {
	model.Params
	// 门店信息
	_stores []Store
}

// NewTaobaoJstAstrolabeStoreinventoryItemqueryRequest 初始化TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest对象
func NewTaobaoJstAstrolabeStoreinventoryItemqueryRequest() *TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest {
	return &TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.itemquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStores is Stores Setter
// 门店信息
func (r *TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// GetStores Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest) GetStores() []Store {
	return r._stores
}
