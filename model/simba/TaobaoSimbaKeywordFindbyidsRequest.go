package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）根据一堆关键词ids获取关键词 APIRequest
taobao.simba.keyword.findbyids

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordFindbyidsRequest struct {
    model.Params

    // 关键词ids
    bidwordIds   []int64 

}

func NewTaobaoSimbaKeywordFindbyidsRequest() *TaobaoSimbaKeywordFindbyidsRequest{
    return &TaobaoSimbaKeywordFindbyidsRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordFindbyidsRequest) GetApiMethodName() string {
    return "taobao.simba.keyword.findbyids"
}

func (r TaobaoSimbaKeywordFindbyidsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordFindbyidsRequest) SetBidwordIds(bidwordIds []int64) error {
    r.bidwordIds = bidwordIds
    r.Set("bidword_ids", bidwordIds)
    return nil
}

func (r TaobaoSimbaKeywordFindbyidsRequest) GetBidwordIds() []int64 {
    return r.bidwordIds
}

