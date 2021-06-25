package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词按流量细分的数据 APIRequest
taobao.simba.insight.wordssubdata.get

获取关键词按流量进行细分的数据，返回结果中network表示流量的来源，意义如下：1->PC站内，2->PC站外,4->无线站内 5->无线站外
*/
type TaobaoSimbaInsightWordssubdataGetRequest struct {
    model.Params

    // 关键词列表
    bidwordList   []String 

    // 开始时间
    startDate   string 

    // 结束时间
    endDate   string 

}

func NewTaobaoSimbaInsightWordssubdataGetRequest() *TaobaoSimbaInsightWordssubdataGetRequest{
    return &TaobaoSimbaInsightWordssubdataGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaInsightWordssubdataGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.wordssubdata.get"
}

func (r TaobaoSimbaInsightWordssubdataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaInsightWordssubdataGetRequest) SetBidwordList(bidwordList []String) error {
    r.bidwordList = bidwordList
    r.Set("bidword_list", bidwordList)
    return nil
}

func (r TaobaoSimbaInsightWordssubdataGetRequest) GetBidwordList() []String {
    return r.bidwordList
}

func (r *TaobaoSimbaInsightWordssubdataGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoSimbaInsightWordssubdataGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoSimbaInsightWordssubdataGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoSimbaInsightWordssubdataGetRequest) GetEndDate() string {
    return r.endDate
}

