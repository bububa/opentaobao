package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest 库存增量更新接口 API请求
// taobao.jst.astrolabe.storeinventory.itemupdate
//
// ERP调用该接口，增量更新门店或电商仓库存，该接口一次可以同时增量更新多个门店的多个商品的非确定性库存。
type TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest struct {
	model.Params
	// 门店列表
	_stores []Store
	// 操作时间
	_operationTime string
}

// NewTaobaoJstAstrolabeStoreinventoryItemupdateRequest 初始化TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest对象
func NewTaobaoJstAstrolabeStoreinventoryItemupdateRequest() *TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest {
	return &TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) Reset() {
	r._stores = r._stores[:0]
	r._operationTime = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.itemupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStores is Stores Setter
// 门店列表
func (r *TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// GetStores Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) GetStores() []Store {
	return r._stores
}

// SetOperationTime is OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) SetOperationTime(_operationTime string) error {
	r._operationTime = _operationTime
	r.Set("operation_time", _operationTime)
	return nil
}

// GetOperationTime OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) GetOperationTime() string {
	return r._operationTime
}

var poolTaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstAstrolabeStoreinventoryItemupdateRequest()
	},
}

// GetTaobaoJstAstrolabeStoreinventoryItemupdateRequest 从 sync.Pool 获取 TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest
func GetTaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest() *TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest {
	return poolTaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest.Get().(*TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest 将 TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest(v *TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest) {
	v.Reset()
	poolTaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest.Put(v)
}
