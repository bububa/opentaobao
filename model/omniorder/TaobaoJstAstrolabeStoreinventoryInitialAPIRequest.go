package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryInitialAPIRequest 后端商品库存初始化 API请求
// taobao.jst.astrolabe.storeinventory.initial
//
// 初始化电商仓或门店库存，该接口一次可以初始化多个门店(或电商仓)的多个商品的多种类型库存。此接口只能使用一次，后续所有的库存变动都需走增量库存同步接口。
type TaobaoJstAstrolabeStoreinventoryInitialAPIRequest struct {
	model.Params
	// 门店列表
	_stores []Store
	// 操作时间
	_operationTime string
}

// NewTaobaoJstAstrolabeStoreinventoryInitialRequest 初始化TaobaoJstAstrolabeStoreinventoryInitialAPIRequest对象
func NewTaobaoJstAstrolabeStoreinventoryInitialRequest() *TaobaoJstAstrolabeStoreinventoryInitialAPIRequest {
	return &TaobaoJstAstrolabeStoreinventoryInitialAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstAstrolabeStoreinventoryInitialAPIRequest) Reset() {
	r._stores = r._stores[:0]
	r._operationTime = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryInitialAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.initial"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryInitialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstAstrolabeStoreinventoryInitialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStores is Stores Setter
// 门店列表
func (r *TaobaoJstAstrolabeStoreinventoryInitialAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// GetStores Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryInitialAPIRequest) GetStores() []Store {
	return r._stores
}

// SetOperationTime is OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryInitialAPIRequest) SetOperationTime(_operationTime string) error {
	r._operationTime = _operationTime
	r.Set("operation_time", _operationTime)
	return nil
}

// GetOperationTime OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryInitialAPIRequest) GetOperationTime() string {
	return r._operationTime
}

var poolTaobaoJstAstrolabeStoreinventoryInitialAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstAstrolabeStoreinventoryInitialRequest()
	},
}

// GetTaobaoJstAstrolabeStoreinventoryInitialRequest 从 sync.Pool 获取 TaobaoJstAstrolabeStoreinventoryInitialAPIRequest
func GetTaobaoJstAstrolabeStoreinventoryInitialAPIRequest() *TaobaoJstAstrolabeStoreinventoryInitialAPIRequest {
	return poolTaobaoJstAstrolabeStoreinventoryInitialAPIRequest.Get().(*TaobaoJstAstrolabeStoreinventoryInitialAPIRequest)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryInitialAPIRequest 将 TaobaoJstAstrolabeStoreinventoryInitialAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstAstrolabeStoreinventoryInitialAPIRequest(v *TaobaoJstAstrolabeStoreinventoryInitialAPIRequest) {
	v.Reset()
	poolTaobaoJstAstrolabeStoreinventoryInitialAPIRequest.Put(v)
}
