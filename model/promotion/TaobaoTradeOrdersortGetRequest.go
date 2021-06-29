package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取前N有礼活动的开奖订单列表 API请求
taobao.trade.ordersort.get

获取前N有礼活动的开奖订单列表
*/
type TaobaoTradeOrdersortGetRequest struct {
    model.Params
    // 活动ID
    _activityId   int64
    // 页码
    _pageNo   int64
    // 一页记录数, 必须写死500
    _pageSize   int64
}

// 初始化TaobaoTradeOrdersortGetRequest对象
func NewTaobaoTradeOrdersortGetRequest() *TaobaoTradeOrdersortGetRequest{
    return &TaobaoTradeOrdersortGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeOrdersortGetRequest) GetApiMethodName() string {
    return "taobao.trade.ordersort.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeOrdersortGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动ID
func (r *TaobaoTradeOrdersortGetRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoTradeOrdersortGetRequest) GetActivityId() int64 {
    return r._activityId
}
// PageNo Setter
// 页码
func (r *TaobaoTradeOrdersortGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTradeOrdersortGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 一页记录数, 必须写死500
func (r *TaobaoTradeOrdersortGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTradeOrdersortGetRequest) GetPageSize() int64 {
    return r._pageSize
}
