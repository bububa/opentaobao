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
    bidwordList   []string
    // 开始时间
    startDate   string
    // 结束时间
    endDate   string
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
func (r *TaobaoSimbaInsightWordssubdataGetRequest) SetBidwordList(bidwordList []string) error {
    r.bidwordList = bidwordList
    r.Set("bidword_list", bidwordList)
    return nil
}

// BidwordList Getter
func (r TaobaoSimbaInsightWordssubdataGetRequest) GetBidwordList() []string {
    return r.bidwordList
}
// StartDate Setter
// 开始时间
func (r *TaobaoSimbaInsightWordssubdataGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSimbaInsightWordssubdataGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 结束时间
func (r *TaobaoSimbaInsightWordssubdataGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSimbaInsightWordssubdataGetRequest) GetEndDate() string {
    return r.endDate
}
