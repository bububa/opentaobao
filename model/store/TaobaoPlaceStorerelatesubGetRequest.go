package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店和子门店关系查找 API请求
taobao.place.storerelatesub.get

门店和子门店关系查找
*/
type TaobaoPlaceStorerelatesubGetRequest struct {
    model.Params
    // 门店Id
    _storeId   int64
    // 查询语句
    _query   string
    // 第几页
    _pageNo   int64
    // 页大小
    _pageSize   int64
}

// 初始化TaobaoPlaceStorerelatesubGetRequest对象
func NewTaobaoPlaceStorerelatesubGetRequest() *TaobaoPlaceStorerelatesubGetRequest{
    return &TaobaoPlaceStorerelatesubGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStorerelatesubGetRequest) GetApiMethodName() string {
    return "taobao.place.storerelatesub.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStorerelatesubGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店Id
func (r *TaobaoPlaceStorerelatesubGetRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStorerelatesubGetRequest) GetStoreId() int64 {
    return r._storeId
}
// Query Setter
// 查询语句
func (r *TaobaoPlaceStorerelatesubGetRequest) SetQuery(_query string) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r TaobaoPlaceStorerelatesubGetRequest) GetQuery() string {
    return r._query
}
// PageNo Setter
// 第几页
func (r *TaobaoPlaceStorerelatesubGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoPlaceStorerelatesubGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 页大小
func (r *TaobaoPlaceStorerelatesubGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoPlaceStorerelatesubGetRequest) GetPageSize() int64 {
    return r._pageSize
}
