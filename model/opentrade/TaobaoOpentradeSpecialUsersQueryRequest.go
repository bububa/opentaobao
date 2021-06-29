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
    _openUserIds   []string
    // 分页参数，每页大小
    _pageSize   int64
    // 商品ID
    _itemId   int64
    // 商品SKU ID，不存在传0
    _skuId   int64
    // 用户状态
    _status   string
    // 分页参数，当前页，以0开始
    _pageIndex   int64
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
func (r *TaobaoOpentradeSpecialUsersQueryRequest) SetOpenUserIds(_openUserIds []string) error {
    r._openUserIds = _openUserIds
    r.Set("open_user_ids", _openUserIds)
    return nil
}

// OpenUserIds Getter
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetOpenUserIds() []string {
    return r._openUserIds
}
// PageSize Setter
// 分页参数，每页大小
func (r *TaobaoOpentradeSpecialUsersQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// ItemId Setter
// 商品ID
func (r *TaobaoOpentradeSpecialUsersQueryRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetItemId() int64 {
    return r._itemId
}
// SkuId Setter
// 商品SKU ID，不存在传0
func (r *TaobaoOpentradeSpecialUsersQueryRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetSkuId() int64 {
    return r._skuId
}
// Status Setter
// 用户状态
func (r *TaobaoOpentradeSpecialUsersQueryRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetStatus() string {
    return r._status
}
// PageIndex Setter
// 分页参数，当前页，以0开始
func (r *TaobaoOpentradeSpecialUsersQueryRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoOpentradeSpecialUsersQueryRequest) GetPageIndex() int64 {
    return r._pageIndex
}
