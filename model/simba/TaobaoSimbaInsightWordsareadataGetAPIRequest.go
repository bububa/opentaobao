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
type TaobaoSimbaInsightWordsareadataGetAPIRequest struct {
    model.Params
    // 关键词
    _bidword   string
    // 开始时间，格式：yyyy-MM-dd
    _startDate   string
    // 结束时间，格式：yyyy-MM-dd
    _endDate   string
}

// 初始化TaobaoSimbaInsightWordsareadataGetAPIRequest对象
func NewTaobaoSimbaInsightWordsareadataGetRequest() *TaobaoSimbaInsightWordsareadataGetAPIRequest{
    return &TaobaoSimbaInsightWordsareadataGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightWordsareadataGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.insight.wordsareadata.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightWordsareadataGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bidword Setter
// 关键词
func (r *TaobaoSimbaInsightWordsareadataGetAPIRequest) SetBidword(_bidword string) error {
    r._bidword = _bidword
    r.Set("bidword", _bidword)
    return nil
}

// Bidword Getter
func (r TaobaoSimbaInsightWordsareadataGetAPIRequest) GetBidword() string {
    return r._bidword
}
// StartDate Setter
// 开始时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightWordsareadataGetAPIRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSimbaInsightWordsareadataGetAPIRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 结束时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightWordsareadataGetAPIRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSimbaInsightWordsareadataGetAPIRequest) GetEndDate() string {
    return r._endDate
}
