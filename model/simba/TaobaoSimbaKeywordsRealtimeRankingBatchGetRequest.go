package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词的新版实时排名 APIRequest
taobao.simba.keywords.realtime.ranking.batch.get

根据关键词ID获取关键词的新版实时排名
*/
type TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest struct {
    model.Params

    // 旺旺名称
    nick   string 

    // adgroupId
    adGroupId   int64 

    // 关键词列表集合,id用半角逗号分割，一次最多20个
    bidwordIds   []int64 

}

func NewTaobaoSimbaKeywordsRealtimeRankingBatchGetRequest() *TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest{
    return &TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywords.realtime.ranking.batch.get"
}

func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) SetAdGroupId(adGroupId int64) error {
    r.adGroupId = adGroupId
    r.Set("ad_group_id", adGroupId)
    return nil
}

func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) GetAdGroupId() int64 {
    return r.adGroupId
}

func (r *TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) SetBidwordIds(bidwordIds []int64) error {
    r.bidwordIds = bidwordIds
    r.Set("bidword_ids", bidwordIds)
    return nil
}

func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) GetBidwordIds() []int64 {
    return r.bidwordIds
}

