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
type TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest struct {
    model.Params
    // 旺旺名称
    _nick   string
    // adgroupId
    _adGroupId   int64
    // 关键词列表集合,id用半角逗号分割，一次最多20个
    _bidwordIds   []int64
}

// 初始化TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest对象
func NewTaobaoSimbaKeywordsRealtimeRankingBatchGetRequest() *TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest{
    return &TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.keywords.realtime.ranking.batch.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 旺旺名称
func (r *TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest) GetNick() string {
    return r._nick
}
// AdGroupId Setter
// adgroupId
func (r *TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest) SetAdGroupId(_adGroupId int64) error {
    r._adGroupId = _adGroupId
    r.Set("ad_group_id", _adGroupId)
    return nil
}

// AdGroupId Getter
func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest) GetAdGroupId() int64 {
    return r._adGroupId
}
// BidwordIds Setter
// 关键词列表集合,id用半角逗号分割，一次最多20个
func (r *TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest) SetBidwordIds(_bidwordIds []int64) error {
    r._bidwordIds = _bidwordIds
    r.Set("bidword_ids", _bidwordIds)
    return nil
}

// BidwordIds Getter
func (r TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest) GetBidwordIds() []int64 {
    return r._bidwordIds
}
