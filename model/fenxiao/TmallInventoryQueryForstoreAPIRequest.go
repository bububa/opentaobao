package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallInventoryQueryForstoreAPIRequest
查询后端商品仓库库存 API请求
tmall.inventory.query.forstore

商家查询后端商品仓库库存 */
type TmallInventoryQueryForstoreAPIRequest struct {
	model.Params
	// 查询列表
	_paramList []InventoryQueryForStoreRequest
}

// NewTmallInventoryQueryForstoreRequest 初始化TmallInventoryQueryForstoreAPIRequest对象
func NewTmallInventoryQueryForstoreRequest() *TmallInventoryQueryForstoreAPIRequest {
	return &TmallInventoryQueryForstoreAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallInventoryQueryForstoreAPIRequest) GetApiMethodName() string {
	return "tmall.inventory.query.forstore"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallInventoryQueryForstoreAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamList Setter
// 查询列表
func (r *TmallInventoryQueryForstoreAPIRequest) SetParamList(_paramList []InventoryQueryForStoreRequest) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// Get ParamList Getter
func (r TmallInventoryQueryForstoreAPIRequest) GetParamList() []InventoryQueryForStoreRequest {
	return r._paramList
}
