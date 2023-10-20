package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstastrolabestoreinventoryitemqueryAPIRequest 库存查询接口 API请求
// taobao.jst.astrolabe.storeinventory.itemquery
//
// 查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
type TaobaojstastrolabestoreinventoryitemqueryAPIRequest struct {
	model.Params
	// 门店信息
	_stores []Store
}

// NewTaobaojstastrolabestoreinventoryitemqueryRequest 初始化TaobaojstastrolabestoreinventoryitemqueryAPIRequest对象
func NewTaobaojstastrolabestoreinventoryitemqueryRequest() *TaobaojstastrolabestoreinventoryitemqueryAPIRequest {
	return &TaobaojstastrolabestoreinventoryitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstastrolabestoreinventoryitemqueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.itemquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstastrolabestoreinventoryitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstastrolabestoreinventoryitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStores is Stores Setter
// 门店信息
func (r *TaobaojstastrolabestoreinventoryitemqueryAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// GetStores Stores Getter
func (r TaobaojstastrolabestoreinventoryitemqueryAPIRequest) GetStores() []Store {
	return r._stores
}
