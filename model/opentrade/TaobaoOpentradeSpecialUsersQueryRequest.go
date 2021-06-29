package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单标记信息查询 API请求
taobao.opentrade.special.users.query

专属下单标记信息查询
*/
type TaobaoOpentradeSpecialUsersQueryRequest struct {
    model.Params
    // 用户openId列表，多个以逗号(,)分割
    openUserIds   []string
    // 分页参数，每页大小
    pageSize   int64
    // 商品ID
    itemId   int64
    // 商品SKU ID，不存在传0
    skuId   int64
    // 用户状态
    status   string
    // 分页参数，当前页，以0开始
    pageIndex   int64
}

// 初始化TaobaoOpentradeSpecialUsersQueryRequest对象
func NewTaobaoOpentradeSpecialUsersQueryRequest() *TaobaoOpentradeSpecialUsersQueryRequest{
    return &TaobaoOpentradeSpecialUsersQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetApiMethodName() string {
    return "taobao.opentrade.special.users.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenUserIds Setter
// 用户openId列表，多个以逗号(,)分割
func (r *TaobaoOpentradeSpecialUsersQueryRequest) SetOpenUserIds(openUserIds []string) error {
    r.openUserIds = openUserIds
    r.Set("open_user_ids", openUserIds)
    return nil
}

// OpenUserIds Getter
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetOpenUserIds() []string {
    return r.openUserIds
}
// PageSize Setter
// 分页参数，每页大小
func (r *TaobaoOpentradeSpecialUsersQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// ItemId Setter
// 商品ID
func (r *TaobaoOpentradeSpecialUsersQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetItemId() int64 {
    return r.itemId
}
// SkuId Setter
// 商品SKU ID，不存在传0
func (r *TaobaoOpentradeSpecialUsersQueryRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

// SkuId Getter
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetSkuId() int64 {
    return r.skuId
}
// Status Setter
// 用户状态
func (r *TaobaoOpentradeSpecialUsersQueryRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetStatus() string {
    return r.status
}
// PageIndex Setter
// 分页参数，当前页，以0开始
func (r *TaobaoOpentradeSpecialUsersQueryRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetPageIndex() int64 {
    return r.pageIndex
}
