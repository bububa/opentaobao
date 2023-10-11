package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCampaignGetAPIRequest 查询单个计划详情 API请求
// taobao.universalbp.campaign.get
//
// 查询单个计划详情信息（不包括报表数据）
type TaobaoUniversalbpCampaignGetAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// campaignQueryVO
	_campaignQueryVO *CampaignQueryVo
}

// NewTaobaoUniversalbpCampaignGetRequest 初始化TaobaoUniversalbpCampaignGetAPIRequest对象
func NewTaobaoUniversalbpCampaignGetRequest() *TaobaoUniversalbpCampaignGetAPIRequest {
	return &TaobaoUniversalbpCampaignGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpCampaignGetAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.campaign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpCampaignGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpCampaignGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpCampaignGetAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpCampaignGetAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetCampaignQueryVO is CampaignQueryVO Setter
// campaignQueryVO
func (r *TaobaoUniversalbpCampaignGetAPIRequest) SetCampaignQueryVO(_campaignQueryVO *CampaignQueryVo) error {
	r._campaignQueryVO = _campaignQueryVO
	r.Set("campaign_query_v_o", _campaignQueryVO)
	return nil
}

// GetCampaignQueryVO CampaignQueryVO Getter
func (r TaobaoUniversalbpCampaignGetAPIRequest) GetCampaignQueryVO() *CampaignQueryVo {
	return r._campaignQueryVO
}
