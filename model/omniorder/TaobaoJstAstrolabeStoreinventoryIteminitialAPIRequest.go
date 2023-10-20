package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstastrolabestoreinventoryiteminitialAPIRequest 库存初始化接口 API请求
// taobao.jst.astrolabe.storeinventory.iteminitial
//
// ERP调用奇门的接口，对门店的库存进行初始化
type TaobaojstastrolabestoreinventoryiteminitialAPIRequest struct {
	model.Params
	// 门店列表
	_stores []Store
	// 操作时间
	_operationTime string
}

// NewTaobaojstastrolabestoreinventoryiteminitialRequest 初始化TaobaojstastrolabestoreinventoryiteminitialAPIRequest对象
func NewTaobaojstastrolabestoreinventoryiteminitialRequest() *TaobaojstastrolabestoreinventoryiteminitialAPIRequest {
	return &TaobaojstastrolabestoreinventoryiteminitialAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstastrolabestoreinventoryiteminitialAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.iteminitial"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstastrolabestoreinventoryiteminitialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstastrolabestoreinventoryiteminitialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStores is Stores Setter
// 门店列表
func (r *TaobaojstastrolabestoreinventoryiteminitialAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// GetStores Stores Getter
func (r TaobaojstastrolabestoreinventoryiteminitialAPIRequest) GetStores() []Store {
	return r._stores
}

// SetOperationTime is OperationTime Setter
// 操作时间
func (r *TaobaojstastrolabestoreinventoryiteminitialAPIRequest) SetOperationTime(_operationTime string) error {
	r._operationTime = _operationTime
	r.Set("operation_time", _operationTime)
	return nil
}

// GetOperationTime OperationTime Getter
func (r TaobaojstastrolabestoreinventoryiteminitialAPIRequest) GetOperationTime() string {
	return r._operationTime
}
