package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestorerelatesubgetAPIRequest 门店和子门店关系查找 API请求
// taobao.place.storerelatesub.get
//
// 门店和子门店关系查找
type TaobaoplacestorerelatesubgetAPIRequest struct {
	model.Params
	// 查询语句
	_query string
	// 门店Id
	_storeId int64
	// 第几页
	_pageNo int64
	// 页大小
	_pageSize int64
}

// NewTaobaoplacestorerelatesubgetRequest 初始化TaobaoplacestorerelatesubgetAPIRequest对象
func NewTaobaoplacestorerelatesubgetRequest() *TaobaoplacestorerelatesubgetAPIRequest {
	return &TaobaoplacestorerelatesubgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestorerelatesubgetAPIRequest) GetApiMethodName() string {
	return "taobao.place.storerelatesub.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestorerelatesubgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestorerelatesubgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询语句
func (r *TaobaoplacestorerelatesubgetAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TaobaoplacestorerelatesubgetAPIRequest) GetQuery() string {
	return r._query
}

// SetStoreId is StoreId Setter
// 门店Id
func (r *TaobaoplacestorerelatesubgetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoplacestorerelatesubgetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetPageNo is PageNo Setter
// 第几页
func (r *TaobaoplacestorerelatesubgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoplacestorerelatesubgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaoplacestorerelatesubgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoplacestorerelatesubgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
