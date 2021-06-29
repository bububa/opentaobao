package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除一批关键词 API请求
taobao.simba.keywords.delete

删除一批关键词
*/
type TaobaoSimbaKeywordsDeleteRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 推广计划Id
    _campaignId   int64
    // 关键词Id数组，最多100个
    _keywordIds   []int64
}

// 初始化TaobaoSimbaKeywordsDeleteRequest对象
func NewTaobaoSimbaKeywordsDeleteRequest() *TaobaoSimbaKeywordsDeleteRequest{
    return &TaobaoSimbaKeywordsDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.keywords.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsDeleteRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaKeywordsDeleteRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaKeywordsDeleteRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaKeywordsDeleteRequest) GetCampaignId() int64 {
    return r._campaignId
}
// KeywordIds Setter
// 关键词Id数组，最多100个
func (r *TaobaoSimbaKeywordsDeleteRequest) SetKeywordIds(_keywordIds []int64) error {
    r._keywordIds = _keywordIds
    r.Set("keyword_ids", _keywordIds)
    return nil
}

// KeywordIds Getter
func (r TaobaoSimbaKeywordsDeleteRequest) GetKeywordIds() []int64 {
    return r._keywordIds
}
