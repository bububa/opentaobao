package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
尖货交易排队信息查询 API请求
taobao.opentrade.queue.query

尖货交易排队信息查询
*/
type TaobaoOpentradeQueueQueryRequest struct {
    model.Params
    // 排队用户状态，新用户为NEW
    status   string
    // 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
    activityId   string
    // 排队商品SKU ID，不存在传0
    skuId   int64
    // 排队商品ID
    itemId   int64
    // 分页参数，每页大小
    pageSize   int64
    // 分页参数，当前页，以0开始
    pageIndex   int64
}

// 初始化TaobaoOpentradeQueueQueryRequest对象
func NewTaobaoOpentradeQueueQueryRequest() *TaobaoOpentradeQueueQueryRequest{
    return &TaobaoOpentradeQueueQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeQueueQueryRequest) GetApiMethodName() string {
    return "taobao.opentrade.queue.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeQueueQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 排队用户状态，新用户为NEW
func (r *TaobaoOpentradeQueueQueryRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoOpentradeQueueQueryRequest) GetStatus() string {
    return r.status
}
// ActivityId Setter
// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
func (r *TaobaoOpentradeQueueQueryRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoOpentradeQueueQueryRequest) GetActivityId() string {
    return r.activityId
}
// SkuId Setter
// 排队商品SKU ID，不存在传0
func (r *TaobaoOpentradeQueueQueryRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

// SkuId Getter
func (r TaobaoOpentradeQueueQueryRequest) GetSkuId() int64 {
    return r.skuId
}
// ItemId Setter
// 排队商品ID
func (r *TaobaoOpentradeQueueQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeQueueQueryRequest) GetItemId() int64 {
    return r.itemId
}
// PageSize Setter
// 分页参数，每页大小
func (r *TaobaoOpentradeQueueQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpentradeQueueQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageIndex Setter
// 分页参数，当前页，以0开始
func (r *TaobaoOpentradeQueueQueryRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoOpentradeQueueQueryRequest) GetPageIndex() int64 {
    return r.pageIndex
}
