package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadtargettaglistrecommendtagAPIRequest 给计划推荐标签 API请求
// alibaba.scbp.ad.target.tag.list.recommend.tag
//
// 给计划推荐标签
type AlibabascbpadtargettaglistrecommendtagAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划ID
	_campaignId int64
	// 推荐实体类
	_recommendQuery *TargetTagRecommendQueryDto
}

// NewAlibabascbpadtargettaglistrecommendtagRequest 初始化AlibabascbpadtargettaglistrecommendtagAPIRequest对象
func NewAlibabascbpadtargettaglistrecommendtagRequest() *AlibabascbpadtargettaglistrecommendtagAPIRequest {
	return &AlibabascbpadtargettaglistrecommendtagAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadtargettaglistrecommendtagAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.target.tag.list.recommend.tag"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadtargettaglistrecommendtagAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadtargettaglistrecommendtagAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadtargettaglistrecommendtagAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadtargettaglistrecommendtagAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划ID
func (r *AlibabascbpadtargettaglistrecommendtagAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadtargettaglistrecommendtagAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetRecommendQuery is RecommendQuery Setter
// 推荐实体类
func (r *AlibabascbpadtargettaglistrecommendtagAPIRequest) SetRecommendQuery(_recommendQuery *TargetTagRecommendQueryDto) error {
	r._recommendQuery = _recommendQuery
	r.Set("recommend_query", _recommendQuery)
	return nil
}

// GetRecommendQuery RecommendQuery Getter
func (r AlibabascbpadtargettaglistrecommendtagAPIRequest) GetRecommendQuery() *TargetTagRecommendQueryDto {
	return r._recommendQuery
}
