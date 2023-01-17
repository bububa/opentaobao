package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest 后端商品库存增量更新接口 API请求
// taobao.jst.astrolabe.storeinventory.update
//
// 增量更新门店或电商仓库存，该接口一次可以同时增量更新多个门店的多个商品的非确定性库存
type TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest struct {
	model.Params
	// 门店列表
	_stores []Store
	// 操作时间
	_operationTime string
}

// NewTaobaoJstAstrolabeStoreinventoryUpdateRequest 初始化TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest对象
func NewTaobaoJstAstrolabeStoreinventoryUpdateRequest() *TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest {
	return &TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStores is Stores Setter
// 门店列表
func (r *TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest) SetStores(_stores []Store) error {
	r._stores = _stores
	r.Set("stores", _stores)
	return nil
}

// GetStores Stores Getter
func (r TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest) GetStores() []Store {
	return r._stores
}

// SetOperationTime is OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest) SetOperationTime(_operationTime string) error {
	r._operationTime = _operationTime
	r.Set("operation_time", _operationTime)
	return nil
}

// GetOperationTime OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest) GetOperationTime() string {
	return r._operationTime
}
