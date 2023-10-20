package omniorder

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstAstrolabeStoreinventoryQueryAPIRequest) Reset() {
	r._stores = r._stores[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstAstrolabeStoreinventoryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoJstAstrolabeStoreinventoryQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstAstrolabeStoreinventoryQueryRequest()
	},
}

// GetTaobaoJstAstrolabeStoreinventoryQueryRequest 从 sync.Pool 获取 TaobaoJstAstrolabeStoreinventoryQueryAPIRequest
func GetTaobaoJstAstrolabeStoreinventoryQueryAPIRequest() *TaobaoJstAstrolabeStoreinventoryQueryAPIRequest {
	return poolTaobaoJstAstrolabeStoreinventoryQueryAPIRequest.Get().(*TaobaoJstAstrolabeStoreinventoryQueryAPIRequest)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryQueryAPIRequest 将 TaobaoJstAstrolabeStoreinventoryQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstAstrolabeStoreinventoryQueryAPIRequest(v *TaobaoJstAstrolabeStoreinventoryQueryAPIRequest) {
	v.Reset()
	poolTaobaoJstAstrolabeStoreinventoryQueryAPIRequest.Put(v)
}
