package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadgroupcountadgroupAPIRequest 统计adgroup数量 API请求
// alibaba.scbp.ad.group.count.ad.group
//
// 统计adgroup数量
type AlibabascbpadgroupcountadgroupAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 查询条件
	_adGroupQuery *AdGroupQueryDto
}

// NewAlibabascbpadgroupcountadgroupRequest 初始化AlibabascbpadgroupcountadgroupAPIRequest对象
func NewAlibabascbpadgroupcountadgroupRequest() *AlibabascbpadgroupcountadgroupAPIRequest {
	return &AlibabascbpadgroupcountadgroupAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadgroupcountadgroupAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.count.ad.group"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadgroupcountadgroupAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadgroupcountadgroupAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadgroupcountadgroupAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadgroupcountadgroupAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadgroupcountadgroupAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadgroupcountadgroupAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdGroupQuery is AdGroupQuery Setter
// 查询条件
func (r *AlibabascbpadgroupcountadgroupAPIRequest) SetAdGroupQuery(_adGroupQuery *AdGroupQueryDto) error {
	r._adGroupQuery = _adGroupQuery
	r.Set("ad_group_query", _adGroupQuery)
	return nil
}

// GetAdGroupQuery AdGroupQuery Getter
func (r AlibabascbpadgroupcountadgroupAPIRequest) GetAdGroupQuery() *AdGroupQueryDto {
	return r._adGroupQuery
}
