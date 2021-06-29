package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词的新版实时排名 API请求
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

// 初始化TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest对象
func NewTaobaoSimbaKeywordsRealtimeRankingBatchGetRequest() *TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest{
    return &TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywords.realtime.ranking.batch.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 旺旺名称
func (r *TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) GetNick() string {
    return r.nick
}
// AdGroupId Setter
// adgroupId
func (r *TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) SetAdGroupId(adGroupId int64) error {
    r.adGroupId = adGroupId
    r.Set("ad_group_id", adGroupId)
    return nil
}

// AdGroupId Getter
func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) GetAdGroupId() int64 {
    return r.adGroupId
}
// BidwordIds Setter
// 关键词列表集合,id用半角逗号分割，一次最多20个
func (r *TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) SetBidwordIds(bidwordIds []int64) error {
    r.bidwordIds = bidwordIds
    r.Set("bidword_ids", bidwordIds)
    return nil
}

// BidwordIds Getter
func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetRequest) GetBidwordIds() []int64 {
    return r.bidwordIds
}
