package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店关联商品查询接口 API请求
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

// 初始化TaobaoQimenStoreitemQueryRequest对象
func NewTaobaoQimenStoreitemQueryRequest() *TaobaoQimenStoreitemQueryRequest{
    return &TaobaoQimenStoreitemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreitemQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.storeitem.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Page Setter
// 当前页面
func (r *TaobaoQimenStoreitemQueryRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r TaobaoQimenStoreitemQueryRequest) GetPage() int64 {
    return r.page
}
// StoreId Setter
// 线上门店id
func (r *TaobaoQimenStoreitemQueryRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoQimenStoreitemQueryRequest) GetStoreId() int64 {
    return r.storeId
}
