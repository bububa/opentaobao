package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallInventoryQueryForstoreAPIRequest 查询后端商品仓库库存 API请求
// tmall.inventory.query.forstore
//
// 商家查询后端商品仓库库存
type TmallInventoryQueryForstoreAPIRequest struct {
	model.Params
	// 查询列表
	_paramList []InventoryQueryForStoreRequest
}

// NewTmallInventoryQueryForstoreRequest 初始化TmallInventoryQueryForstoreAPIRequest对象
func NewTmallInventoryQueryForstoreRequest() *TmallInventoryQueryForstoreAPIRequest {
	return &TmallInventoryQueryForstoreAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallInventoryQueryForstoreAPIRequest) Reset() {
	r._paramList = r._paramList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallInventoryQueryForstoreAPIRequest) GetApiMethodName() string {
	return "tmall.inventory.query.forstore"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallInventoryQueryForstoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallInventoryQueryForstoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 查询列表
func (r *TmallInventoryQueryForstoreAPIRequest) SetParamList(_paramList []InventoryQueryForStoreRequest) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r TmallInventoryQueryForstoreAPIRequest) GetParamList() []InventoryQueryForStoreRequest {
	return r._paramList
}

var poolTmallInventoryQueryForstoreAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallInventoryQueryForstoreRequest()
	},
}

// GetTmallInventoryQueryForstoreRequest 从 sync.Pool 获取 TmallInventoryQueryForstoreAPIRequest
func GetTmallInventoryQueryForstoreAPIRequest() *TmallInventoryQueryForstoreAPIRequest {
	return poolTmallInventoryQueryForstoreAPIRequest.Get().(*TmallInventoryQueryForstoreAPIRequest)
}

// ReleaseTmallInventoryQueryForstoreAPIRequest 将 TmallInventoryQueryForstoreAPIRequest 放入 sync.Pool
func ReleaseTmallInventoryQueryForstoreAPIRequest(v *TmallInventoryQueryForstoreAPIRequest) {
	v.Reset()
	poolTmallInventoryQueryForstoreAPIRequest.Put(v)
}
