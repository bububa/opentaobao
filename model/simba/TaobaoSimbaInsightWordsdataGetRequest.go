package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词的大盘数据 APIRequest
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

func NewTaobaoSimbaInsightWordsdataGetRequest() *TaobaoSimbaInsightWordsdataGetRequest{
    return &TaobaoSimbaInsightWordsdataGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaInsightWordsdataGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.wordsdata.get"
}

func (r TaobaoSimbaInsightWordsdataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaInsightWordsdataGetRequest) SetBidwordList(bidwordList []string) error {
    r.bidwordList = bidwordList
    r.Set("bidword_list", bidwordList)
    return nil
}

func (r TaobaoSimbaInsightWordsdataGetRequest) GetBidwordList() []string {
    return r.bidwordList
}

func (r *TaobaoSimbaInsightWordsdataGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoSimbaInsightWordsdataGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoSimbaInsightWordsdataGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoSimbaInsightWordsdataGetRequest) GetEndDate() string {
    return r.endDate
}

