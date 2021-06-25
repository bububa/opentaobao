package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词按地域进行细分的数据 APIRequest
taobao.simba.insight.wordsareadata.get

获取关键词按地域细分的详细数据，目前地域只能细化到省级别，返回的结果中包括市，是为了方便以后扩展，目前结果中市的值等于省。
*/
type TaobaoSimbaInsightWordsareadataGetRequest struct {
    model.Params

    // 关键词
    bidword   string 

    // 开始时间，格式：yyyy-MM-dd
    startDate   string 

    // 结束时间，格式：yyyy-MM-dd
    endDate   string 

}

func NewTaobaoSimbaInsightWordsareadataGetRequest() *TaobaoSimbaInsightWordsareadataGetRequest{
    return &TaobaoSimbaInsightWordsareadataGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaInsightWordsareadataGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.wordsareadata.get"
}

func (r TaobaoSimbaInsightWordsareadataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaInsightWordsareadataGetRequest) SetBidword(bidword string) error {
    r.bidword = bidword
    r.Set("bidword", bidword)
    return nil
}

func (r TaobaoSimbaInsightWordsareadataGetRequest) GetBidword() string {
    return r.bidword
}

func (r *TaobaoSimbaInsightWordsareadataGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoSimbaInsightWordsareadataGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoSimbaInsightWordsareadataGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoSimbaInsightWordsareadataGetRequest) GetEndDate() string {
    return r.endDate
}

