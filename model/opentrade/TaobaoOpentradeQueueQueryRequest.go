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
    _status   string
    // 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
    _activityId   string
    // 排队商品SKU ID，不存在传0
    _skuId   int64
    // 排队商品ID
    _itemId   int64
    // 分页参数，每页大小
    _pageSize   int64
    // 分页参数，当前页，以0开始
    _pageIndex   int64
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
func (r *TaobaoOpentradeQueueQueryRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoOpentradeQueueQueryRequest) GetStatus() string {
    return r._status
}
// ActivityId Setter
// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
func (r *TaobaoOpentradeQueueQueryRequest) SetActivityId(_activityId string) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoOpentradeQueueQueryRequest) GetActivityId() string {
    return r._activityId
}
// SkuId Setter
// 排队商品SKU ID，不存在传0
func (r *TaobaoOpentradeQueueQueryRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoOpentradeQueueQueryRequest) GetSkuId() int64 {
    return r._skuId
}
// ItemId Setter
// 排队商品ID
func (r *TaobaoOpentradeQueueQueryRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeQueueQueryRequest) GetItemId() int64 {
    return r._itemId
}
// PageSize Setter
// 分页参数，每页大小
func (r *TaobaoOpentradeQueueQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpentradeQueueQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageIndex Setter
// 分页参数，当前页，以0开始
func (r *TaobaoOpentradeQueueQueryRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoOpentradeQueueQueryRequest) GetPageIndex() int64 {
    return r._pageIndex
}
