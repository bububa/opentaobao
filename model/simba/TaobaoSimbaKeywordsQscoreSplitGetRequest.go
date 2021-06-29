package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新质量分服务 API请求
taobao.simba.keywords.qscore.split.get

获取关键词新的质量分
*/
type TaobaoSimbaKeywordsQscoreSplitGetRequest struct {
    model.Params
    // 账号昵称
    nick   string
    // 推广组id
    adGroupId   int64
    // 词id数组（最多批量获取20个）
    bidwordIds   []int64
}

// 初始化TaobaoSimbaKeywordsQscoreSplitGetRequest对象
func NewTaobaoSimbaKeywordsQscoreSplitGetRequest() *TaobaoSimbaKeywordsQscoreSplitGetRequest{
    return &TaobaoSimbaKeywordsQscoreSplitGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsQscoreSplitGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywords.qscore.split.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsQscoreSplitGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 账号昵称
func (r *TaobaoSimbaKeywordsQscoreSplitGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaKeywordsQscoreSplitGetRequest) GetNick() string {
    return r.nick
}
// AdGroupId Setter
// 推广组id
func (r *TaobaoSimbaKeywordsQscoreSplitGetRequest) SetAdGroupId(adGroupId int64) error {
    r.adGroupId = adGroupId
    r.Set("ad_group_id", adGroupId)
    return nil
}

// AdGroupId Getter
func (r TaobaoSimbaKeywordsQscoreSplitGetRequest) GetAdGroupId() int64 {
    return r.adGroupId
}
// BidwordIds Setter
// 词id数组（最多批量获取20个）
func (r *TaobaoSimbaKeywordsQscoreSplitGetRequest) SetBidwordIds(bidwordIds []int64) error {
    r.bidwordIds = bidwordIds
    r.Set("bidword_ids", bidwordIds)
    return nil
}

// BidwordIds Getter
func (r TaobaoSimbaKeywordsQscoreSplitGetRequest) GetBidwordIds() []int64 {
    return r.bidwordIds
}
