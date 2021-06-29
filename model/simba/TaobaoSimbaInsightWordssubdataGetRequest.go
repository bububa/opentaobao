package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词按流量细分的数据 API请求
taobao.simba.insight.wordssubdata.get

获取关键词按流量进行细分的数据，返回结果中network表示流量的来源，意义如下：1->PC站内，2->PC站外,4->无线站内 5->无线站外
*/
type TaobaoSimbaInsightWordssubdataGetRequest struct {
    model.Params
    // 关键词列表
    _bidwordList   []string
    // 开始时间
    _startDate   string
    // 结束时间
    _endDate   string
}

// 初始化TaobaoSimbaInsightWordssubdataGetRequest对象
func NewTaobaoSimbaInsightWordssubdataGetRequest() *TaobaoSimbaInsightWordssubdataGetRequest{
    return &TaobaoSimbaInsightWordssubdataGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightWordssubdataGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.wordssubdata.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightWordssubdataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BidwordList Setter
// 关键词列表
func (r *TaobaoSimbaInsightWordssubdataGetRequest) SetBidwordList(_bidwordList []string) error {
    r._bidwordList = _bidwordList
    r.Set("bidword_list", _bidwordList)
    return nil
}

// BidwordList Getter
func (r TaobaoSimbaInsightWordssubdataGetRequest) GetBidwordList() []string {
    return r._bidwordList
}
// StartDate Setter
// 开始时间
func (r *TaobaoSimbaInsightWordssubdataGetRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSimbaInsightWordssubdataGetRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 结束时间
func (r *TaobaoSimbaInsightWordssubdataGetRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSimbaInsightWordssubdataGetRequest) GetEndDate() string {
    return r._endDate
}
