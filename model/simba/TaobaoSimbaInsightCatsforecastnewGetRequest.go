package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取词的相关类目预测数据 APIRequest
taobao.simba.insight.catsforecastnew.get

根据给定的词，预测这些词的相关类目
*/
type TaobaoSimbaInsightCatsforecastnewGetRequest struct {
    model.Params

    // 需要查询的词列表
    bidwordList   []string 

}

func NewTaobaoSimbaInsightCatsforecastnewGetRequest() *TaobaoSimbaInsightCatsforecastnewGetRequest{
    return &TaobaoSimbaInsightCatsforecastnewGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaInsightCatsforecastnewGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.catsforecastnew.get"
}

func (r TaobaoSimbaInsightCatsforecastnewGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaInsightCatsforecastnewGetRequest) SetBidwordList(bidwordList []string) error {
    r.bidwordList = bidwordList
    r.Set("bidword_list", bidwordList)
    return nil
}

func (r TaobaoSimbaInsightCatsforecastnewGetRequest) GetBidwordList() []string {
    return r.bidwordList
}

