package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】订单搜索 APIRequest
taobao.jipiao.agent.order.search

卖家根据条件查询淘宝订单id列表
*/
type TaobaoJipiaoAgentOrderSearchRequest struct {
    model.Params

    // 创建订单时间范围的开始时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认开始时间为当前时间往前推三天 （具体天数可能调整）
    beginTime   string 

    // 创建订单时间范围的结束时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认为当前时间 （具体天数可能调整）
    endTime   string 

    // 订单状态，默认为空，查询所有状态的订单<br/>1:待确认<br/>2:待出票<br/>3:强制成功<br/>4:未付款<br/>5:预订成功<br/>6:已失效
    status   int64 

    // 航程类型： 0.单程；1.往返
    tripType   int64 

    // 是否需要行程单，true表示需要行程单；false表示不许要
    hasItinerary   bool 

    // 页码，默认第一页；注意：页大小固定，搜索结果中返回页大小pageSize，和是否包含下一页hasNext
    page   int64 

    // 扩展字段:<br/>needFilterAutobook：默认true。待出票状态下，会根据此值过滤掉自动出票状态下订单，以防止重复出票的问题。对于精选票，此值需要设置成false，并由API使用者保证不会重复出票。
    extra   string 

}

func NewTaobaoJipiaoAgentOrderSearchRequest() *TaobaoJipiaoAgentOrderSearchRequest{
    return &TaobaoJipiaoAgentOrderSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJipiaoAgentOrderSearchRequest) GetApiMethodName() string {
    return "taobao.jipiao.agent.order.search"
}

func (r TaobaoJipiaoAgentOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJipiaoAgentOrderSearchRequest) SetBeginTime(beginTime string) error {
    r.beginTime = beginTime
    r.Set("begin_time", beginTime)
    return nil
}

func (r TaobaoJipiaoAgentOrderSearchRequest) GetBeginTime() string {
    return r.beginTime
}

func (r *TaobaoJipiaoAgentOrderSearchRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoJipiaoAgentOrderSearchRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoJipiaoAgentOrderSearchRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoJipiaoAgentOrderSearchRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoJipiaoAgentOrderSearchRequest) SetTripType(tripType int64) error {
    r.tripType = tripType
    r.Set("trip_type", tripType)
    return nil
}

func (r TaobaoJipiaoAgentOrderSearchRequest) GetTripType() int64 {
    return r.tripType
}

func (r *TaobaoJipiaoAgentOrderSearchRequest) SetHasItinerary(hasItinerary bool) error {
    r.hasItinerary = hasItinerary
    r.Set("has_itinerary", hasItinerary)
    return nil
}

func (r TaobaoJipiaoAgentOrderSearchRequest) GetHasItinerary() bool {
    return r.hasItinerary
}

func (r *TaobaoJipiaoAgentOrderSearchRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r TaobaoJipiaoAgentOrderSearchRequest) GetPage() int64 {
    return r.page
}

func (r *TaobaoJipiaoAgentOrderSearchRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

func (r TaobaoJipiaoAgentOrderSearchRequest) GetExtra() string {
    return r.extra
}

