package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstastrolabestoreinventoryitemupdateAPIRequest 库存增量更新接口 API请求
// taobao.jst.astrolabe.storeinventory.itemupdate
//
// ERP调用该接口，增量更新门店或电商仓库存，该接口一次可以同时增量更新多个门店的多个商品的非确定性库存。
type TaobaojstastrolabestoreinventoryitemupdateAPIRequest struct {
	model.Params
	// 门店列表
	_stores []Store
	// 操作时间
	_operationTime string
}

// NewTaobaojstastrolabestoreinventoryitemupdateRequest 初始化TaobaojstastrolabestoreinventoryitemupdateAPIRequest对象
func NewTaobaojstastrolabestoreinventoryitemupdateRequest() *TaobaojstastrolabestoreinventoryitemupdateAPIRequest {
	return &TaobaojstastrolabestoreinventoryitemupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstastrolabestoreinventoryitemupdateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.itemupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstastrolabestoreinventoryitemupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstastrolabestoreinventoryitemupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStores is Stores Setter
// 门店列表
func (r *TaobaojstastrolabestoreinventoryitemupdateAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// GetStores Stores Getter
func (r TaobaojstastrolabestoreinventoryitemupdateAPIRequest) GetStores() []Store {
	return r._stores
}

// SetOperationTime is OperationTime Setter
// 操作时间
func (r *TaobaojstastrolabestoreinventoryitemupdateAPIRequest) SetOperationTime(_operationTime string) error {
	r._operationTime = _operationTime
	r.Set("operation_time", _operationTime)
	return nil
}

// GetOperationTime OperationTime Getter
func (r TaobaojstastrolabestoreinventoryitemupdateAPIRequest) GetOperationTime() string {
	return r._operationTime
}
