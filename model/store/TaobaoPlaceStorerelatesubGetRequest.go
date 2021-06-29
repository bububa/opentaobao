package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店和子门店关系查找 APIRequest
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

func NewTaobaoPlaceStorerelatesubGetRequest() *TaobaoPlaceStorerelatesubGetRequest{
    return &TaobaoPlaceStorerelatesubGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStorerelatesubGetRequest) GetApiMethodName() string {
    return "taobao.place.storerelatesub.get"
}

func (r TaobaoPlaceStorerelatesubGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStorerelatesubGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoPlaceStorerelatesubGetRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoPlaceStorerelatesubGetRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r TaobaoPlaceStorerelatesubGetRequest) GetQuery() string {
    return r.query
}

func (r *TaobaoPlaceStorerelatesubGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoPlaceStorerelatesubGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoPlaceStorerelatesubGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoPlaceStorerelatesubGetRequest) GetPageSize() int64 {
    return r.pageSize
}

