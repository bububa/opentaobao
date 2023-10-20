package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryStoreQueryAPIRequest 查询仓库信息 API请求
// taobao.inventory.store.query
//
// 查询商家仓信息
type TaobaoInventoryStoreQueryAPIRequest struct {
	model.Params
	// 商家的仓库编码
	_storeCode string
}

// NewTaobaoInventoryStoreQueryRequest 初始化TaobaoInventoryStoreQueryAPIRequest对象
func NewTaobaoInventoryStoreQueryRequest() *TaobaoInventoryStoreQueryAPIRequest {
	return &TaobaoInventoryStoreQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoInventoryStoreQueryAPIRequest) Reset() {
	r._storeCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryStoreQueryAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.store.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryStoreQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoInventoryStoreQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreCode is StoreCode Setter
// 商家的仓库编码
func (r *TaobaoInventoryStoreQueryAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoInventoryStoreQueryAPIRequest) GetStoreCode() string {
	return r._storeCode
}

var poolTaobaoInventoryStoreQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoInventoryStoreQueryRequest()
	},
}

// GetTaobaoInventoryStoreQueryRequest 从 sync.Pool 获取 TaobaoInventoryStoreQueryAPIRequest
func GetTaobaoInventoryStoreQueryAPIRequest() *TaobaoInventoryStoreQueryAPIRequest {
	return poolTaobaoInventoryStoreQueryAPIRequest.Get().(*TaobaoInventoryStoreQueryAPIRequest)
}

// ReleaseTaobaoInventoryStoreQueryAPIRequest 将 TaobaoInventoryStoreQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoInventoryStoreQueryAPIRequest(v *TaobaoInventoryStoreQueryAPIRequest) {
	v.Reset()
	poolTaobaoInventoryStoreQueryAPIRequest.Put(v)
}
