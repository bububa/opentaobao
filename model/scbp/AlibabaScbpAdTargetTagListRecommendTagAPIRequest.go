package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdTargetTagListRecommendTagAPIRequest 给计划推荐标签 API请求
// alibaba.scbp.ad.target.tag.list.recommend.tag
//
// 给计划推荐标签
type AlibabaScbpAdTargetTagListRecommendTagAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划ID
	_campaignId int64
	// 推荐实体类
	_recommendQuery *TargetTagRecommendQueryDto
}

// NewAlibabaScbpAdTargetTagListRecommendTagRequest 初始化AlibabaScbpAdTargetTagListRecommendTagAPIRequest对象
func NewAlibabaScbpAdTargetTagListRecommendTagRequest() *AlibabaScbpAdTargetTagListRecommendTagAPIRequest {
	return &AlibabaScbpAdTargetTagListRecommendTagAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdTargetTagListRecommendTagAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.target.tag.list.recommend.tag"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdTargetTagListRecommendTagAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdTargetTagListRecommendTagAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdTargetTagListRecommendTagAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdTargetTagListRecommendTagAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划ID
func (r *AlibabaScbpAdTargetTagListRecommendTagAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdTargetTagListRecommendTagAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetRecommendQuery is RecommendQuery Setter
// 推荐实体类
func (r *AlibabaScbpAdTargetTagListRecommendTagAPIRequest) SetRecommendQuery(_recommendQuery *TargetTagRecommendQueryDto) error {
	r._recommendQuery = _recommendQuery
	r.Set("recommend_query", _recommendQuery)
	return nil
}

// GetRecommendQuery RecommendQuery Getter
func (r AlibabaScbpAdTargetTagListRecommendTagAPIRequest) GetRecommendQuery() *TargetTagRecommendQueryDto {
	return r._recommendQuery
}
