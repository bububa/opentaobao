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
type TaobaoSimbaInsightWordsdataGetRequest struct {
    model.Params
    // 关键词列表，最多可传100个。
    bidwordList   []string
    // 开始时间
    startDate   string
    // 结束时间
    endDate   string
}

// 初始化TaobaoSimbaInsightWordsdataGetRequest对象
func NewTaobaoSimbaInsightWordsdataGetRequest() *TaobaoSimbaInsightWordsdataGetRequest{
    return &TaobaoSimbaInsightWordsdataGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightWordsdataGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.wordsdata.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightWordsdataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BidwordList Setter
// 关键词列表，最多可传100个。
func (r *TaobaoSimbaInsightWordsdataGetRequest) SetBidwordList(bidwordList []string) error {
    r.bidwordList = bidwordList
    r.Set("bidword_list", bidwordList)
    return nil
}

// BidwordList Getter
func (r TaobaoSimbaInsightWordsdataGetRequest) GetBidwordList() []string {
    return r.bidwordList
}
// StartDate Setter
// 开始时间
func (r *TaobaoSimbaInsightWordsdataGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSimbaInsightWordsdataGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 结束时间
func (r *TaobaoSimbaInsightWordsdataGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSimbaInsightWordsdataGetRequest) GetEndDate() string {
    return r.endDate
}
