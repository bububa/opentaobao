package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest 库存初始化接口 API请求
// taobao.jst.astrolabe.storeinventory.iteminitial
//
// ERP调用奇门的接口，对门店的库存进行初始化
type TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest struct {
	model.Params
	// 门店列表
	_stores []Store
	// 操作时间
	_operationTime string
}

// NewTaobaoJstAstrolabeStoreinventoryIteminitialRequest 初始化TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest对象
func NewTaobaoJstAstrolabeStoreinventoryIteminitialRequest() *TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest {
	return &TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) Reset() {
	r._stores = r._stores[:0]
	r._operationTime = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.iteminitial"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStores is Stores Setter
// 门店列表
func (r *TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// GetStores Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) GetStores() []Store {
	return r._stores
}

// SetOperationTime is OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) SetOperationTime(_operationTime string) error {
	r._operationTime = _operationTime
	r.Set("operation_time", _operationTime)
	return nil
}

// GetOperationTime OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) GetOperationTime() string {
	return r._operationTime
}

var poolTaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstAstrolabeStoreinventoryIteminitialRequest()
	},
}

// GetTaobaoJstAstrolabeStoreinventoryIteminitialRequest 从 sync.Pool 获取 TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest
func GetTaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest() *TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest {
	return poolTaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest.Get().(*TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest 将 TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest(v *TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) {
	v.Reset()
	poolTaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest.Put(v)
}
