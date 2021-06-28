package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)销量明星质量分相关接口 APIRequest
taobao.simba.salestar.keywords.qscore.split.get

获取关键词新的质量分
*/
type TaobaoSimbaSalestarKeywordsQscoreSplitGetRequest struct {
    model.Params

    // 账号昵称
    nick   string 

    // 推广组id
    adGroupId   int64 

    // 词id数组（最多批量获取20个）
    bidwordIds   []int64 

}

func NewTaobaoSimbaSalestarKeywordsQscoreSplitGetRequest() *TaobaoSimbaSalestarKeywordsQscoreSplitGetRequest{
    return &TaobaoSimbaSalestarKeywordsQscoreSplitGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSalestarKeywordsQscoreSplitGetRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.keywords.qscore.split.get"
}

func (r TaobaoSimbaSalestarKeywordsQscoreSplitGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSalestarKeywordsQscoreSplitGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaSalestarKeywordsQscoreSplitGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaSalestarKeywordsQscoreSplitGetRequest) SetAdGroupId(adGroupId int64) error {
    r.adGroupId = adGroupId
    r.Set("ad_group_id", adGroupId)
    return nil
}

func (r TaobaoSimbaSalestarKeywordsQscoreSplitGetRequest) GetAdGroupId() int64 {
    return r.adGroupId
}

func (r *TaobaoSimbaSalestarKeywordsQscoreSplitGetRequest) SetBidwordIds(bidwordIds []int64) error {
    r.bidwordIds = bidwordIds
    r.Set("bidword_ids", bidwordIds)
    return nil
}

func (r TaobaoSimbaSalestarKeywordsQscoreSplitGetRequest) GetBidwordIds() []int64 {
    return r.bidwordIds
}

