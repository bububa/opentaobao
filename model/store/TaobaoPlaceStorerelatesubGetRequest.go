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
    storeId   int64
    // 查询语句
    query   string
    // 第几页
    pageNo   int64
    // 页大小
    pageSize   int64
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
func (r *TaobaoPlaceStorerelatesubGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStorerelatesubGetRequest) GetStoreId() int64 {
    return r.storeId
}
// Query Setter
// 查询语句
func (r *TaobaoPlaceStorerelatesubGetRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r TaobaoPlaceStorerelatesubGetRequest) GetQuery() string {
    return r.query
}
// PageNo Setter
// 第几页
func (r *TaobaoPlaceStorerelatesubGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoPlaceStorerelatesubGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 页大小
func (r *TaobaoPlaceStorerelatesubGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoPlaceStorerelatesubGetRequest) GetPageSize() int64 {
    return r.pageSize
}
