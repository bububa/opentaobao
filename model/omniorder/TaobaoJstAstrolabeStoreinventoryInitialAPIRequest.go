package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstastrolabestoreinventoryinitialAPIRequest 后端商品库存初始化 API请求
// taobao.jst.astrolabe.storeinventory.initial
//
// 初始化电商仓或门店库存，该接口一次可以初始化多个门店(或电商仓)的多个商品的多种类型库存。此接口只能使用一次，后续所有的库存变动都需走增量库存同步接口。
type TaobaojstastrolabestoreinventoryinitialAPIRequest struct {
	model.Params
	// 门店列表
	_stores []Store
	// 操作时间
	_operationTime string
}

// NewTaobaojstastrolabestoreinventoryinitialRequest 初始化TaobaojstastrolabestoreinventoryinitialAPIRequest对象
func NewTaobaojstastrolabestoreinventoryinitialRequest() *TaobaojstastrolabestoreinventoryinitialAPIRequest {
	return &TaobaojstastrolabestoreinventoryinitialAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstastrolabestoreinventoryinitialAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.initial"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstastrolabestoreinventoryinitialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstastrolabestoreinventoryinitialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStores is Stores Setter
// 门店列表
func (r *TaobaojstastrolabestoreinventoryinitialAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// GetStores Stores Getter
func (r TaobaojstastrolabestoreinventoryinitialAPIRequest) GetStores() []Store {
	return r._stores
}

// SetOperationTime is OperationTime Setter
// 操作时间
func (r *TaobaojstastrolabestoreinventoryinitialAPIRequest) SetOperationTime(_operationTime string) error {
	r._operationTime = _operationTime
	r.Set("operation_time", _operationTime)
	return nil
}

// GetOperationTime OperationTime Getter
func (r TaobaojstastrolabestoreinventoryinitialAPIRequest) GetOperationTime() string {
	return r._operationTime
}
