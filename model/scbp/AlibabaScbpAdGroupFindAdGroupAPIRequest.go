package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadgroupfindadgroupAPIRequest 查询推广组 API请求
// alibaba.scbp.ad.group.find.ad.group
//
// 查询推广组
type AlibabascbpadgroupfindadgroupAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 入参
	_adGroupQuery *AdGroupQueryDto
}

// NewAlibabascbpadgroupfindadgroupRequest 初始化AlibabascbpadgroupfindadgroupAPIRequest对象
func NewAlibabascbpadgroupfindadgroupRequest() *AlibabascbpadgroupfindadgroupAPIRequest {
	return &AlibabascbpadgroupfindadgroupAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadgroupfindadgroupAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.find.ad.group"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadgroupfindadgroupAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadgroupfindadgroupAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadgroupfindadgroupAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadgroupfindadgroupAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadgroupfindadgroupAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadgroupfindadgroupAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdGroupQuery is AdGroupQuery Setter
// 入参
func (r *AlibabascbpadgroupfindadgroupAPIRequest) SetAdGroupQuery(_adGroupQuery *AdGroupQueryDto) error {
	r._adGroupQuery = _adGroupQuery
	r.Set("ad_group_query", _adGroupQuery)
	return nil
}

// GetAdGroupQuery AdGroupQuery Getter
func (r AlibabascbpadgroupfindadgroupAPIRequest) GetAdGroupQuery() *AdGroupQueryDto {
	return r._adGroupQuery
}
