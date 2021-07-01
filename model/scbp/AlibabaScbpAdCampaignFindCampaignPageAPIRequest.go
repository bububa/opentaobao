package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdCampaignFindCampaignPageAPIRequest
分页查询计划 API请求
alibaba.scbp.ad.campaign.find.campaign.page

分页查询计划 */
type AlibabaScbpAdCampaignFindCampaignPageAPIRequest struct {
	model.Params
	// 请求实体类
	_campaignQuery *CampaignQueryDto
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdCampaignFindCampaignPageRequest 初始化AlibabaScbpAdCampaignFindCampaignPageAPIRequest对象
func NewAlibabaScbpAdCampaignFindCampaignPageRequest() *AlibabaScbpAdCampaignFindCampaignPageAPIRequest {
	return &AlibabaScbpAdCampaignFindCampaignPageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignFindCampaignPageAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.find.campaign.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignFindCampaignPageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampaignQuery Setter
// 请求实体类
func (r *AlibabaScbpAdCampaignFindCampaignPageAPIRequest) SetCampaignQuery(_campaignQuery *CampaignQueryDto) error {
	r._campaignQuery = _campaignQuery
	r.Set("campaign_query", _campaignQuery)
	return nil
}

// Get CampaignQuery Getter
func (r AlibabaScbpAdCampaignFindCampaignPageAPIRequest) GetCampaignQuery() *CampaignQueryDto {
	return r._campaignQuery
}

// Set is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignFindCampaignPageAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// Get TopContext Getter
func (r AlibabaScbpAdCampaignFindCampaignPageAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
