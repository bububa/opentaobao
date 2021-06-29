package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取词的相关词 API请求
taobao.simba.insight.relatedwords.get

获取给定词的若干相关词，返回结果中越相关的权重越大，排在越前面，根据number参数对返回结果进行截断。
*/
type TaobaoSimbaInsightRelatedwordsGetRequest struct {
    model.Params
    // 要查询的词列表
    _bidwordList   []string
    // 表示返回数据的条数
    _number   int64
}

// 初始化TaobaoSimbaInsightRelatedwordsGetRequest对象
func NewTaobaoSimbaInsightRelatedwordsGetRequest() *TaobaoSimbaInsightRelatedwordsGetRequest{
    return &TaobaoSimbaInsightRelatedwordsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightRelatedwordsGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.relatedwords.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightRelatedwordsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BidwordList Setter
// 要查询的词列表
func (r *TaobaoSimbaInsightRelatedwordsGetRequest) SetBidwordList(_bidwordList []string) error {
    r._bidwordList = _bidwordList
    r.Set("bidword_list", _bidwordList)
    return nil
}

// BidwordList Getter
func (r TaobaoSimbaInsightRelatedwordsGetRequest) GetBidwordList() []string {
    return r._bidwordList
}
// Number Setter
// 表示返回数据的条数
func (r *TaobaoSimbaInsightRelatedwordsGetRequest) SetNumber(_number int64) error {
    r._number = _number
    r.Set("number", _number)
    return nil
}

// Number Getter
func (r TaobaoSimbaInsightRelatedwordsGetRequest) GetNumber() int64 {
    return r._number
}
