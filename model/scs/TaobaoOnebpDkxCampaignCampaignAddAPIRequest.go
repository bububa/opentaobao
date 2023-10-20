package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxcampaigncampaignaddAPIRequest 添加计划 API请求
// taobao.onebp.dkx.campaign.campaign.add
//
// 添加计划
// 套餐包计划： {&#34;campaign_view&#34;:{&#34;ad_strategy_info&#34;:{&#34;budget&#34;:&#34;2000&#34;,&#34;item_ids&#34;:[617764134830],&#34;launch_strategy_type&#34;:&#34;cost_first&#34;,&#34;order_amount&#34;:&#34;2000&#34;},&#34;campaign_name&#34;:&#34;新建byTop测试&#34;,&#34;launch_time&#34;:{&#34;begin_time&#34;:&#34;2021-09-27 00:00:01&#34;,&#34;end_time&#34;:&#34;2021-09-31 00:00:01&#34;},&#34;marketing&#34;:{&#34;market_aim&#34;:1041,&#34;market_scene&#34;:1095},&#34;promotion_model&#34;:&#34;order&#34;}}
// 持续推广计划：
// {&#34;campaignViewDTO&#34;:{&#34;adStrategyInfo&#34;:{&#34;hide&#34;:false,&#34;itemIds&#34;:[45031793073],&#34;launchStrategyType&#34;:&#34;cost_first&#34;},&#34;campaignName&#34;:&#34;新建byTop&#34;,&#34;dayBudget&#34;:{&#34;dayBudget&#34;:&#34;100&#34;,&#34;dmcType&#34;:1},&#34;marketing&#34;:{&#34;marketAim&#34;:1037,&#34;marketScene&#34;:1023},&#34;promotionModel&#34;:&#34;daily&#34;}}
type TaobaoonebpdkxcampaigncampaignaddAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 创建计划入参
	_solutionResult *SolutionTopDto
}

// NewTaobaoonebpdkxcampaigncampaignaddRequest 初始化TaobaoonebpdkxcampaigncampaignaddAPIRequest对象
func NewTaobaoonebpdkxcampaigncampaignaddRequest() *TaobaoonebpdkxcampaigncampaignaddAPIRequest {
	return &TaobaoonebpdkxcampaigncampaignaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxcampaigncampaignaddAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.campaign.campaign.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxcampaigncampaignaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxcampaigncampaignaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoonebpdkxcampaigncampaignaddAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoonebpdkxcampaigncampaignaddAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetSolutionResult is SolutionResult Setter
// 创建计划入参
func (r *TaobaoonebpdkxcampaigncampaignaddAPIRequest) SetSolutionResult(_solutionResult *SolutionTopDto) error {
	r._solutionResult = _solutionResult
	r.Set("solution_result", _solutionResult)
	return nil
}

// GetSolutionResult SolutionResult Getter
func (r TaobaoonebpdkxcampaigncampaignaddAPIRequest) GetSolutionResult() *SolutionTopDto {
	return r._solutionResult
}
