package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销量明星关键词删除 APIRequest
taobao.simba.salestar.keywords.delete

（新）关键词删除相关接口
*/
type TaobaoSimbaSalestarKeywordsDeleteRequest struct {
    model.Params

    // 关键词ids
    bidwordIds   []Number 

}

func NewTaobaoSimbaSalestarKeywordsDeleteRequest() *TaobaoSimbaSalestarKeywordsDeleteRequest{
    return &TaobaoSimbaSalestarKeywordsDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSalestarKeywordsDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.keywords.delete"
}

func (r TaobaoSimbaSalestarKeywordsDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSalestarKeywordsDeleteRequest) SetBidwordIds(bidwordIds []Number) error {
    r.bidwordIds = bidwordIds
    r.Set("bidword_ids", bidwordIds)
    return nil
}

func (r TaobaoSimbaSalestarKeywordsDeleteRequest) GetBidwordIds() []Number {
    return r.bidwordIds
}

