package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）关键词更新相关接口 APIRequest
taobao.simba.keyword.update

（新）关键词更新相关接口
*/
type TaobaoSimbaKeywordUpdateRequest struct {
    model.Params

    // 关键词相关信息
    bidwords   []SiriusBidwordDto 

}

func NewTaobaoSimbaKeywordUpdateRequest() *TaobaoSimbaKeywordUpdateRequest{
    return &TaobaoSimbaKeywordUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.keyword.update"
}

func (r TaobaoSimbaKeywordUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordUpdateRequest) SetBidwords(bidwords []SiriusBidwordDto) error {
    r.bidwords = bidwords
    r.Set("bidwords", bidwords)
    return nil
}

func (r TaobaoSimbaKeywordUpdateRequest) GetBidwords() []SiriusBidwordDto {
    return r.bidwords
}

