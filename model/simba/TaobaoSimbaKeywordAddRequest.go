package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）关键词新增接口 APIRequest
taobao.simba.keyword.add

（新）关键词更新相关接口
*/
type TaobaoSimbaKeywordAddRequest struct {
    model.Params

    // 关键词相关信息
    bidwords   []SiriusBidwordDto 

    // 推广单元id
    adgroupId   int64 

}

func NewTaobaoSimbaKeywordAddRequest() *TaobaoSimbaKeywordAddRequest{
    return &TaobaoSimbaKeywordAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordAddRequest) GetApiMethodName() string {
    return "taobao.simba.keyword.add"
}

func (r TaobaoSimbaKeywordAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordAddRequest) SetBidwords(bidwords []SiriusBidwordDto) error {
    r.bidwords = bidwords
    r.Set("bidwords", bidwords)
    return nil
}

func (r TaobaoSimbaKeywordAddRequest) GetBidwords() []SiriusBidwordDto {
    return r.bidwords
}

func (r *TaobaoSimbaKeywordAddRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaKeywordAddRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

