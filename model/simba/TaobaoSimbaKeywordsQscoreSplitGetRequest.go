package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新质量分服务 APIRequest
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
    bidwordIds   []Number 

}

func NewTaobaoSimbaKeywordsQscoreSplitGetRequest() *TaobaoSimbaKeywordsQscoreSplitGetRequest{
    return &TaobaoSimbaKeywordsQscoreSplitGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordsQscoreSplitGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywords.qscore.split.get"
}

func (r TaobaoSimbaKeywordsQscoreSplitGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordsQscoreSplitGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaKeywordsQscoreSplitGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaKeywordsQscoreSplitGetRequest) SetAdGroupId(adGroupId int64) error {
    r.adGroupId = adGroupId
    r.Set("ad_group_id", adGroupId)
    return nil
}

func (r TaobaoSimbaKeywordsQscoreSplitGetRequest) GetAdGroupId() int64 {
    return r.adGroupId
}

func (r *TaobaoSimbaKeywordsQscoreSplitGetRequest) SetBidwordIds(bidwordIds []Number) error {
    r.bidwordIds = bidwordIds
    r.Set("bidword_ids", bidwordIds)
    return nil
}

func (r TaobaoSimbaKeywordsQscoreSplitGetRequest) GetBidwordIds() []Number {
    return r.bidwordIds
}

