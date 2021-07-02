package omniorder

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.iteminitial"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Stores Setter
// 门店列表
func (r *TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// Get Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) GetStores() []Store {
	return r._stores
}

// Set is OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) SetOperationTime(_operationTime string) error {
	r._operationTime = _operationTime
	r.Set("operation_time", _operationTime)
	return nil
}

// Get OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest) GetOperationTime() string {
	return r._operationTime
}
