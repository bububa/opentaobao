package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstastrolabestoreinventoryqueryAPIRequest 后端商品库存查询接口 API请求
// taobao.jst.astrolabe.storeinventory.query
//
// 查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
type TaobaojstastrolabestoreinventoryqueryAPIRequest struct {
	model.Params
	// 门店
	_stores []Store
}

// NewTaobaojstastrolabestoreinventoryqueryRequest 初始化TaobaojstastrolabestoreinventoryqueryAPIRequest对象
func NewTaobaojstastrolabestoreinventoryqueryRequest() *TaobaojstastrolabestoreinventoryqueryAPIRequest {
	return &TaobaojstastrolabestoreinventoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstastrolabestoreinventoryqueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstastrolabestoreinventoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstastrolabestoreinventoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStores is Stores Setter
// 门店
func (r *TaobaojstastrolabestoreinventoryqueryAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// GetStores Stores Getter
func (r TaobaojstastrolabestoreinventoryqueryAPIRequest) GetStores() []Store {
	return r._stores
}
