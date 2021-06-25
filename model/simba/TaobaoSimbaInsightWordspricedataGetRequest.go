package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词按竞价区间的分布数据 APIRequest
taobao.simba.insight.wordspricedata.get

获取关键词按竞价区间进行细分的数据
*/
type TaobaoSimbaInsightWordspricedataGetRequest struct {
    model.Params

    // 关键词
    bidword   string 

    // 开始时间，格式：yyyy-MM-dd
    startDate   string 

    // 结束时间，格式：yyyy-MM-dd
    endDate   string 

}

func NewTaobaoSimbaInsightWordspricedataGetRequest() *TaobaoSimbaInsightWordspricedataGetRequest{
    return &TaobaoSimbaInsightWordspricedataGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaInsightWordspricedataGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.wordspricedata.get"
}

func (r TaobaoSimbaInsightWordspricedataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaInsightWordspricedataGetRequest) SetBidword(bidword string) error {
    r.bidword = bidword
    r.Set("bidword", bidword)
    return nil
}

func (r TaobaoSimbaInsightWordspricedataGetRequest) GetBidword() string {
    return r.bidword
}

func (r *TaobaoSimbaInsightWordspricedataGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoSimbaInsightWordspricedataGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoSimbaInsightWordspricedataGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoSimbaInsightWordspricedataGetRequest) GetEndDate() string {
    return r.endDate
}

