package nrt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStoreRelationQueryAPIRequest 喵零门店关系查询 API请求
// tmall.nrt.store.relation.query
//
// 喵零门店关系查询
type TmallNrtStoreRelationQueryAPIRequest struct {
	model.Params
	// 查询对象
	_storeQuery *StoreQuery
}

// NewTmallNrtStoreRelationQueryRequest 初始化TmallNrtStoreRelationQueryAPIRequest对象
func NewTmallNrtStoreRelationQueryRequest() *TmallNrtStoreRelationQueryAPIRequest {
	return &TmallNrtStoreRelationQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtStoreRelationQueryAPIRequest) Reset() {
	r._storeQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtStoreRelationQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.store.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtStoreRelationQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtStoreRelationQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreQuery is StoreQuery Setter
// 查询对象
func (r *TmallNrtStoreRelationQueryAPIRequest) SetStoreQuery(_storeQuery *StoreQuery) error {
	r._storeQuery = _storeQuery
	r.Set("store_query", _storeQuery)
	return nil
}

// GetStoreQuery StoreQuery Getter
func (r TmallNrtStoreRelationQueryAPIRequest) GetStoreQuery() *StoreQuery {
	return r._storeQuery
}

var poolTmallNrtStoreRelationQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtStoreRelationQueryRequest()
	},
}

// GetTmallNrtStoreRelationQueryRequest 从 sync.Pool 获取 TmallNrtStoreRelationQueryAPIRequest
func GetTmallNrtStoreRelationQueryAPIRequest() *TmallNrtStoreRelationQueryAPIRequest {
	return poolTmallNrtStoreRelationQueryAPIRequest.Get().(*TmallNrtStoreRelationQueryAPIRequest)
}

// ReleaseTmallNrtStoreRelationQueryAPIRequest 将 TmallNrtStoreRelationQueryAPIRequest 放入 sync.Pool
func ReleaseTmallNrtStoreRelationQueryAPIRequest(v *TmallNrtStoreRelationQueryAPIRequest) {
	v.Reset()
	poolTmallNrtStoreRelationQueryAPIRequest.Put(v)
}
