package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取词的相关类目预测数据 API请求
taobao.simba.insight.catsforecastnew.get

根据给定的词，预测这些词的相关类目
*/
type TaobaoSimbaInsightCatsforecastnewGetRequest struct {
    model.Params
    // 需要查询的词列表
    bidwordList   []string
}

// 初始化TaobaoSimbaInsightCatsforecastnewGetRequest对象
func NewTaobaoSimbaInsightCatsforecastnewGetRequest() *TaobaoSimbaInsightCatsforecastnewGetRequest{
    return &TaobaoSimbaInsightCatsforecastnewGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightCatsforecastnewGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.catsforecastnew.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightCatsforecastnewGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BidwordList Setter
// 需要查询的词列表
func (r *TaobaoSimbaInsightCatsforecastnewGetRequest) SetBidwordList(bidwordList []string) error {
    r.bidwordList = bidwordList
    r.Set("bidword_list", bidwordList)
    return nil
}

// BidwordList Getter
func (r TaobaoSimbaInsightCatsforecastnewGetRequest) GetBidwordList() []string {
    return r.bidwordList
}
