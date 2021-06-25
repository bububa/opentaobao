package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取词的相关词 APIRequest
taobao.simba.insight.relatedwords.get

获取给定词的若干相关词，返回结果中越相关的权重越大，排在越前面，根据number参数对返回结果进行截断。
*/
type TaobaoSimbaInsightRelatedwordsGetRequest struct {
    model.Params

    // 要查询的词列表
    bidwordList   []String 

    // 表示返回数据的条数
    number   int64 

}

func NewTaobaoSimbaInsightRelatedwordsGetRequest() *TaobaoSimbaInsightRelatedwordsGetRequest{
    return &TaobaoSimbaInsightRelatedwordsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaInsightRelatedwordsGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.relatedwords.get"
}

func (r TaobaoSimbaInsightRelatedwordsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaInsightRelatedwordsGetRequest) SetBidwordList(bidwordList []String) error {
    r.bidwordList = bidwordList
    r.Set("bidword_list", bidwordList)
    return nil
}

func (r TaobaoSimbaInsightRelatedwordsGetRequest) GetBidwordList() []String {
    return r.bidwordList
}

func (r *TaobaoSimbaInsightRelatedwordsGetRequest) SetNumber(number int64) error {
    r.number = number
    r.Set("number", number)
    return nil
}

func (r TaobaoSimbaInsightRelatedwordsGetRequest) GetNumber() int64 {
    return r.number
}

