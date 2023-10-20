package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxcampaigncampaignmodifyAPIRequest 修改计划 API请求
// taobao.onebp.dkx.campaign.campaign.modify
//
// 修改计划。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoonebpdkxcampaigncampaignmodifyAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 计划查询参数
	_campaignQuery *CampaignQueryTopDto
}

// NewTaobaoonebpdkxcampaigncampaignmodifyRequest 初始化TaobaoonebpdkxcampaigncampaignmodifyAPIRequest对象
func NewTaobaoonebpdkxcampaigncampaignmodifyRequest() *TaobaoonebpdkxcampaigncampaignmodifyAPIRequest {
	return &TaobaoonebpdkxcampaigncampaignmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxcampaigncampaignmodifyAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.campaign.campaign.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxcampaigncampaignmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxcampaigncampaignmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoonebpdkxcampaigncampaignmodifyAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoonebpdkxcampaigncampaignmodifyAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetCampaignQuery is CampaignQuery Setter
// 计划查询参数
func (r *TaobaoonebpdkxcampaigncampaignmodifyAPIRequest) SetCampaignQuery(_campaignQuery *CampaignQueryTopDto) error {
	r._campaignQuery = _campaignQuery
	r.Set("campaign_query", _campaignQuery)
	return nil
}

// GetCampaignQuery CampaignQuery Getter
func (r TaobaoonebpdkxcampaigncampaignmodifyAPIRequest) GetCampaignQuery() *CampaignQueryTopDto {
	return r._campaignQuery
}
