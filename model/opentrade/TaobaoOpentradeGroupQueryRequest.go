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
    itemId   int64
    // 用户openId
    openUserId   string
    // 0 返回未成团列表，1 返回已成团列表
    orderBy   int64
    // 页数
    pageIndex   int64
    // 每页展示条数，不能超过100
    pageSize   int64
    // 组团活动id
    groupActivityId   int64
    // 是否返回已过期的团，true 返回，false 不返回
    withExpire   bool
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
func (r *TaobaoOpentradeGroupQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeGroupQueryRequest) GetItemId() int64 {
    return r.itemId
}
// OpenUserId Setter
// 用户openId
func (r *TaobaoOpentradeGroupQueryRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

// OpenUserId Getter
func (r TaobaoOpentradeGroupQueryRequest) GetOpenUserId() string {
    return r.openUserId
}
// OrderBy Setter
// 0 返回未成团列表，1 返回已成团列表
func (r *TaobaoOpentradeGroupQueryRequest) SetOrderBy(orderBy int64) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

// OrderBy Getter
func (r TaobaoOpentradeGroupQueryRequest) GetOrderBy() int64 {
    return r.orderBy
}
// PageIndex Setter
// 页数
func (r *TaobaoOpentradeGroupQueryRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoOpentradeGroupQueryRequest) GetPageIndex() int64 {
    return r.pageIndex
}
// PageSize Setter
// 每页展示条数，不能超过100
func (r *TaobaoOpentradeGroupQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpentradeGroupQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// GroupActivityId Setter
// 组团活动id
func (r *TaobaoOpentradeGroupQueryRequest) SetGroupActivityId(groupActivityId int64) error {
    r.groupActivityId = groupActivityId
    r.Set("group_activity_id", groupActivityId)
    return nil
}

// GroupActivityId Getter
func (r TaobaoOpentradeGroupQueryRequest) GetGroupActivityId() int64 {
    return r.groupActivityId
}
// WithExpire Setter
// 是否返回已过期的团，true 返回，false 不返回
func (r *TaobaoOpentradeGroupQueryRequest) SetWithExpire(withExpire bool) error {
    r.withExpire = withExpire
    r.Set("with_expire", withExpire)
    return nil
}

// WithExpire Getter
func (r TaobaoOpentradeGroupQueryRequest) GetWithExpire() bool {
    return r.withExpire
}
