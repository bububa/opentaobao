package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词的大盘数据 API请求
taobao.simba.insight.wordsdata.get

获取关键词的详细数据
*/
type TaobaoSimbaInsightWordsdataGetAPIRequest struct {
    model.Params
    // 关键词列表，最多可传100个。
    _bidwordList   []string
    // 开始时间
    _startDate   string
    // 结束时间
    _endDate   string
}

// 初始化TaobaoSimbaInsightWordsdataGetAPIRequest对象
func NewTaobaoSimbaInsightWordsdataGetRequest() *TaobaoSimbaInsightWordsdataGetAPIRequest{
    return &TaobaoSimbaInsightWordsdataGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightWordsdataGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.insight.wordsdata.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightWordsdataGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BidwordList Setter
// 关键词列表，最多可传100个。
func (r *TaobaoSimbaInsightWordsdataGetAPIRequest) SetBidwordList(_bidwordList []string) error {
    r._bidwordList = _bidwordList
    r.Set("bidword_list", _bidwordList)
    return nil
}

// BidwordList Getter
func (r TaobaoSimbaInsightWordsdataGetAPIRequest) GetBidwordList() []string {
    return r._bidwordList
}
// StartDate Setter
// 开始时间
func (r *TaobaoSimbaInsightWordsdataGetAPIRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSimbaInsightWordsdataGetAPIRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 结束时间
func (r *TaobaoSimbaInsightWordsdataGetAPIRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSimbaInsightWordsdataGetAPIRequest) GetEndDate() string {
    return r._endDate
}
