package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsDeleteAPIRequest 删除一批关键词 API请求
// taobao.simba.keywords.delete
//
// 删除一批关键词
type TaobaoSimbaKeywordsDeleteAPIRequest struct {
	model.Params
	// 关键词Id数组，最多100个
	_keywordIds []int64
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaoSimbaKeywordsDeleteRequest 初始化TaobaoSimbaKeywordsDeleteAPIRequest对象
func NewTaobaoSimbaKeywordsDeleteRequest() *TaobaoSimbaKeywordsDeleteAPIRequest {
	return &TaobaoSimbaKeywordsDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetKeywordIds is KeywordIds Setter
// 关键词Id数组，最多100个
func (r *TaobaoSimbaKeywordsDeleteAPIRequest) SetKeywordIds(_keywordIds []int64) error {
	r._keywordIds = _keywordIds
	r.Set("keyword_ids", _keywordIds)
	return nil
}

// GetKeywordIds KeywordIds Getter
func (r TaobaoSimbaKeywordsDeleteAPIRequest) GetKeywordIds() []int64 {
	return r._keywordIds
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsDeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaKeywordsDeleteAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaKeywordsDeleteAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaKeywordsDeleteAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
