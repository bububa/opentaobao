package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）根据一堆关键词ids获取关键词 API请求
taobao.simba.keyword.findbyids

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordFindbyidsRequest struct {
    model.Params
    // 关键词ids
    _bidwordIds   []int64
}

// 初始化TaobaoSimbaKeywordFindbyidsRequest对象
func NewTaobaoSimbaKeywordFindbyidsRequest() *TaobaoSimbaKeywordFindbyidsRequest{
    return &TaobaoSimbaKeywordFindbyidsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordFindbyidsRequest) GetApiMethodName() string {
    return "taobao.simba.keyword.findbyids"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordFindbyidsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BidwordIds Setter
// 关键词ids
func (r *TaobaoSimbaKeywordFindbyidsRequest) SetBidwordIds(_bidwordIds []int64) error {
    r._bidwordIds = _bidwordIds
    r.Set("bidword_ids", _bidwordIds)
    return nil
}

// BidwordIds Getter
func (r TaobaoSimbaKeywordFindbyidsRequest) GetBidwordIds() []int64 {
    return r._bidwordIds
}
