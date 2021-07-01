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
type TaobaoSimbaKeywordFindbyidsAPIRequest struct {
    model.Params
    // 关键词ids
    _bidwordIds   []int64
}

// 初始化TaobaoSimbaKeywordFindbyidsAPIRequest对象
func NewTaobaoSimbaKeywordFindbyidsRequest() *TaobaoSimbaKeywordFindbyidsAPIRequest{
    return &TaobaoSimbaKeywordFindbyidsAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordFindbyidsAPIRequest) GetApiMethodName() string {
    return "taobao.simba.keyword.findbyids"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordFindbyidsAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BidwordIds Setter
// 关键词ids
func (r *TaobaoSimbaKeywordFindbyidsAPIRequest) SetBidwordIds(_bidwordIds []int64) error {
    r._bidwordIds = _bidwordIds
    r.Set("bidword_ids", _bidwordIds)
    return nil
}

// BidwordIds Getter
func (r TaobaoSimbaKeywordFindbyidsAPIRequest) GetBidwordIds() []int64 {
    return r._bidwordIds
}
