package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品关联门店查询接口 APIRequest
taobao.qimen.itemstore.query

商家在ERP等系统中调用该接口，查询线上商品所关联的门店列表
*/
type TaobaoQimenItemstoreQueryRequest struct {
    model.Params

    // 当前查询的页面编码
    page   int64 

    // 线上商品ID
    itemId   int64 

}

func NewTaobaoQimenItemstoreQueryRequest() *TaobaoQimenItemstoreQueryRequest{
    return &TaobaoQimenItemstoreQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenItemstoreQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.itemstore.query"
}

func (r TaobaoQimenItemstoreQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenItemstoreQueryRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r TaobaoQimenItemstoreQueryRequest) GetPage() int64 {
    return r.page
}

func (r *TaobaoQimenItemstoreQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoQimenItemstoreQueryRequest) GetItemId() int64 {
    return r.itemId
}

