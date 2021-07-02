package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStorerelatesubGetAPIRequest 门店和子门店关系查找 API请求
// taobao.place.storerelatesub.get
//
// 门店和子门店关系查找
type TaobaoPlaceStorerelatesubGetAPIRequest struct {
	model.Params
	// 门店Id
	_storeId int64
	// 查询语句
	_query string
	// 第几页
	_pageNo int64
	// 页大小
	_pageSize int64
}

// NewTaobaoPlaceStorerelatesubGetRequest 初始化TaobaoPlaceStorerelatesubGetAPIRequest对象
func NewTaobaoPlaceStorerelatesubGetRequest() *TaobaoPlaceStorerelatesubGetAPIRequest {
	return &TaobaoPlaceStorerelatesubGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStorerelatesubGetAPIRequest) GetApiMethodName() string {
	return "taobao.place.storerelatesub.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStorerelatesubGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStoreId is StoreId Setter
// 门店Id
func (r *TaobaoPlaceStorerelatesubGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoPlaceStorerelatesubGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetQuery is Query Setter
// 查询语句
func (r *TaobaoPlaceStorerelatesubGetAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TaobaoPlaceStorerelatesubGetAPIRequest) GetQuery() string {
	return r._query
}

// SetPageNo is PageNo Setter
// 第几页
func (r *TaobaoPlaceStorerelatesubGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoPlaceStorerelatesubGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaoPlaceStorerelatesubGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoPlaceStorerelatesubGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
