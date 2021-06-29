package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词按竞价区间的分布数据 API请求
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

// 初始化TaobaoSimbaInsightWordspricedataGetRequest对象
func NewTaobaoSimbaInsightWordspricedataGetRequest() *TaobaoSimbaInsightWordspricedataGetRequest{
    return &TaobaoSimbaInsightWordspricedataGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightWordspricedataGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.wordspricedata.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightWordspricedataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bidword Setter
// 关键词
func (r *TaobaoSimbaInsightWordspricedataGetRequest) SetBidword(bidword string) error {
    r.bidword = bidword
    r.Set("bidword", bidword)
    return nil
}

// Bidword Getter
func (r TaobaoSimbaInsightWordspricedataGetRequest) GetBidword() string {
    return r.bidword
}
// StartDate Setter
// 开始时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightWordspricedataGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSimbaInsightWordspricedataGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 结束时间，格式：yyyy-MM-dd
func (r *TaobaoSimbaInsightWordspricedataGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSimbaInsightWordspricedataGetRequest) GetEndDate() string {
    return r.endDate
}
