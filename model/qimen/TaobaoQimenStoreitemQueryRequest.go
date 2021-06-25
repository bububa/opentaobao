package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店关联商品查询接口 APIRequest
taobao.qimen.storeitem.query

商家在ERP等系统中调用该接口，查询某门店所关联的线上商品列表
*/
type TaobaoQimenStoreitemQueryRequest struct {
    model.Params

    // 当前页面
    page   int64 

    // 线上门店id
    storeId   int64 

}

func NewTaobaoQimenStoreitemQueryRequest() *TaobaoQimenStoreitemQueryRequest{
    return &TaobaoQimenStoreitemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenStoreitemQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.storeitem.query"
}

func (r TaobaoQimenStoreitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenStoreitemQueryRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r TaobaoQimenStoreitemQueryRequest) GetPage() int64 {
    return r.page
}

func (r *TaobaoQimenStoreitemQueryRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoQimenStoreitemQueryRequest) GetStoreId() int64 {
    return r.storeId
}

