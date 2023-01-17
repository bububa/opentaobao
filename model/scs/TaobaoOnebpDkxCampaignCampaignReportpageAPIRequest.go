package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest 获取场景计划的报表数据 API请求
// taobao.onebp.dkx.campaign.campaign.reportpage
//
// 获取场景计划的报表数据。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 计划查询参数
	_campaignQuery *CampaignQueryTopDto
}

// NewTaobaoOnebpDkxCampaignCampaignReportpageRequest 初始化TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest对象
func NewTaobaoOnebpDkxCampaignCampaignReportpageRequest() *TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest {
	return &TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.campaign.campaign.reportpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetCampaignQuery is CampaignQuery Setter
// 计划查询参数
func (r *TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest) SetCampaignQuery(_campaignQuery *CampaignQueryTopDto) error {
	r._campaignQuery = _campaignQuery
	r.Set("campaign_query", _campaignQuery)
	return nil
}

// GetCampaignQuery CampaignQuery Getter
func (r TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest) GetCampaignQuery() *CampaignQueryTopDto {
	return r._campaignQuery
}
