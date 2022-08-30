package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCampaignCampaignAddAPIRequest 添加计划 API请求
// taobao.onebp.dkx.campaign.campaign.add
//
// 添加计划
// 套餐包计划： {"campaign_view":{"ad_strategy_info":{"budget":"2000","item_ids":[617764134830],"launch_strategy_type":"cost_first","order_amount":"2000"},"campaign_name":"新建byTop测试","launch_time":{"begin_time":"2021-09-27 00:00:01","end_time":"2021-09-31 00:00:01"},"marketing":{"market_aim":1041,"market_scene":1095},"promotion_model":"order"}}
// 持续推广计划：
// {"campaignViewDTO":{"adStrategyInfo":{"hide":false,"itemIds":[45031793073],"launchStrategyType":"cost_first"},"campaignName":"新建byTop","dayBudget":{"dayBudget":"100","dmcType":1},"marketing":{"marketAim":1037,"marketScene":1023},"promotionModel":"daily"}}
type TaobaoOnebpDkxCampaignCampaignAddAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 创建计划入参
	_solutionResult *SolutionTopDto
}

// NewTaobaoOnebpDkxCampaignCampaignAddRequest 初始化TaobaoOnebpDkxCampaignCampaignAddAPIRequest对象
func NewTaobaoOnebpDkxCampaignCampaignAddRequest() *TaobaoOnebpDkxCampaignCampaignAddAPIRequest {
	return &TaobaoOnebpDkxCampaignCampaignAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxCampaignCampaignAddAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.campaign.campaign.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxCampaignCampaignAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxCampaignCampaignAddAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxCampaignCampaignAddAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetSolutionResult is SolutionResult Setter
// 创建计划入参
func (r *TaobaoOnebpDkxCampaignCampaignAddAPIRequest) SetSolutionResult(_solutionResult *SolutionTopDto) error {
	r._solutionResult = _solutionResult
	r.Set("solution_result", _solutionResult)
	return nil
}

// GetSolutionResult SolutionResult Getter
func (r TaobaoOnebpDkxCampaignCampaignAddAPIRequest) GetSolutionResult() *SolutionTopDto {
	return r._solutionResult
}
