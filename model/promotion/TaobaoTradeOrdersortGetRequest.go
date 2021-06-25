package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取前N有礼活动的开奖订单列表 APIRequest
taobao.trade.ordersort.get

获取前N有礼活动的开奖订单列表
*/
type TaobaoTradeOrdersortGetRequest struct {
    model.Params

    // 活动ID
    activityId   int64 

    // 页码
    pageNo   int64 

    // 一页记录数, 必须写死500
    pageSize   int64 

}

func NewTaobaoTradeOrdersortGetRequest() *TaobaoTradeOrdersortGetRequest{
    return &TaobaoTradeOrdersortGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeOrdersortGetRequest) GetApiMethodName() string {
    return "taobao.trade.ordersort.get"
}

func (r TaobaoTradeOrdersortGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeOrdersortGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoTradeOrdersortGetRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *TaobaoTradeOrdersortGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoTradeOrdersortGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoTradeOrdersortGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoTradeOrdersortGetRequest) GetPageSize() int64 {
    return r.pageSize
}

