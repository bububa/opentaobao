package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购场景查询商品组团详情 API请求
taobao.opentrade.group.query

组团购场景下，查询商品开团详情
*/
type TaobaoOpentradeGroupQueryRequest struct {
    model.Params
    // 商品id
    _itemId   int64
    // 用户openId
    _openUserId   string
    // 0 返回未成团列表，1 返回已成团列表
    _orderBy   int64
    // 页数
    _pageIndex   int64
    // 每页展示条数，不能超过100
    _pageSize   int64
    // 组团活动id
    _groupActivityId   int64
    // 是否返回已过期的团，true 返回，false 不返回
    _withExpire   bool
}

// 初始化TaobaoOpentradeGroupQueryRequest对象
func NewTaobaoOpentradeGroupQueryRequest() *TaobaoOpentradeGroupQueryRequest{
    return &TaobaoOpentradeGroupQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupQueryRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoOpentradeGroupQueryRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeGroupQueryRequest) GetItemId() int64 {
    return r._itemId
}
// OpenUserId Setter
// 用户openId
func (r *TaobaoOpentradeGroupQueryRequest) SetOpenUserId(_openUserId string) error {
    r._openUserId = _openUserId
    r.Set("open_user_id", _openUserId)
    return nil
}

// OpenUserId Getter
func (r TaobaoOpentradeGroupQueryRequest) GetOpenUserId() string {
    return r._openUserId
}
// OrderBy Setter
// 0 返回未成团列表，1 返回已成团列表
func (r *TaobaoOpentradeGroupQueryRequest) SetOrderBy(_orderBy int64) error {
    r._orderBy = _orderBy
    r.Set("order_by", _orderBy)
    return nil
}

// OrderBy Getter
func (r TaobaoOpentradeGroupQueryRequest) GetOrderBy() int64 {
    return r._orderBy
}
// PageIndex Setter
// 页数
func (r *TaobaoOpentradeGroupQueryRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoOpentradeGroupQueryRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// PageSize Setter
// 每页展示条数，不能超过100
func (r *TaobaoOpentradeGroupQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpentradeGroupQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// GroupActivityId Setter
// 组团活动id
func (r *TaobaoOpentradeGroupQueryRequest) SetGroupActivityId(_groupActivityId int64) error {
    r._groupActivityId = _groupActivityId
    r.Set("group_activity_id", _groupActivityId)
    return nil
}

// GroupActivityId Getter
func (r TaobaoOpentradeGroupQueryRequest) GetGroupActivityId() int64 {
    return r._groupActivityId
}
// WithExpire Setter
// 是否返回已过期的团，true 返回，false 不返回
func (r *TaobaoOpentradeGroupQueryRequest) SetWithExpire(_withExpire bool) error {
    r._withExpire = _withExpire
    r.Set("with_expire", _withExpire)
    return nil
}

// WithExpire Getter
func (r TaobaoOpentradeGroupQueryRequest) GetWithExpire() bool {
    return r._withExpire
}
