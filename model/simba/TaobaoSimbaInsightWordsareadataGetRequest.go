package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词按地域进行细分的数据 API请求
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

// 初始化TaobaoSimbaInsightWordsareadataGetRequest对象
func NewTaobaoSimbaInsightWordsareadataGetRequest() *TaobaoSimbaInsightWordsareadataGetRequest{
    return &TaobaoSimbaInsightWordsareadataGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightWordsareadataGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.wordsareadata.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightWordsareadataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bidword Setter
// 关键词
func (r *TaobaoSimbaInsightWordsareadataGetRequest) SetBidword(bidword string) error {
    r.bidword = bidword
    r.Set("bidword", bidword)
    return nil
}

// Bidword Getter
func (r TaobaoSimbaInsightWordsareadataGetRequest) GetBidword() string {
    return r.bidword
}
// StartDate Setter
// 开始时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightWordsareadataGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSimbaInsightWordsareadataGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 结束时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightWordsareadataGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSimbaInsightWordsareadataGetRequest) GetEndDate() string {
    return r.endDate
}
