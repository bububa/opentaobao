package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品关联门店查询接口 API请求
taobao.qimen.itemstore.query

商家在ERP等系统中调用该接口，查询线上商品所关联的门店列表
*/
type TaobaoQimenItemstoreQueryAPIRequest struct {
    model.Params
    // 当前查询的页面编码
    _page   int64
    // 线上商品ID
    _itemId   int64
}

// 初始化TaobaoQimenItemstoreQueryAPIRequest对象
func NewTaobaoQimenItemstoreQueryRequest() *TaobaoQimenItemstoreQueryAPIRequest{
    return &TaobaoQimenItemstoreQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemstoreQueryAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.itemstore.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemstoreQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Page Setter
// 当前查询的页面编码
func (r *TaobaoQimenItemstoreQueryAPIRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r TaobaoQimenItemstoreQueryAPIRequest) GetPage() int64 {
    return r._page
}
// ItemId Setter
// 线上商品ID
func (r *TaobaoQimenItemstoreQueryAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoQimenItemstoreQueryAPIRequest) GetItemId() int64 {
    return r._itemId
}
