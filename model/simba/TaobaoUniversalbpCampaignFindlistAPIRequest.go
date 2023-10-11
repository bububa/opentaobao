package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCampaignFindlistAPIRequest 查询全量计划列表(不分页) API请求
// taobao.universalbp.campaign.findlist
//
// 查询场景内的全量计划列表
type TaobaoUniversalbpCampaignFindlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// campaignQueryVO
	_campaignQueryVO *CampaignQueryVo
}

// NewTaobaoUniversalbpCampaignFindlistRequest 初始化TaobaoUniversalbpCampaignFindlistAPIRequest对象
func NewTaobaoUniversalbpCampaignFindlistRequest() *TaobaoUniversalbpCampaignFindlistAPIRequest {
	return &TaobaoUniversalbpCampaignFindlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpCampaignFindlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.campaign.findlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpCampaignFindlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpCampaignFindlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpCampaignFindlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpCampaignFindlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetCampaignQueryVO is CampaignQueryVO Setter
// campaignQueryVO
func (r *TaobaoUniversalbpCampaignFindlistAPIRequest) SetCampaignQueryVO(_campaignQueryVO *CampaignQueryVo) error {
	r._campaignQueryVO = _campaignQueryVO
	r.Set("campaign_query_v_o", _campaignQueryVO)
	return nil
}

// GetCampaignQueryVO CampaignQueryVO Getter
func (r TaobaoUniversalbpCampaignFindlistAPIRequest) GetCampaignQueryVO() *CampaignQueryVo {
	return r._campaignQueryVO
}
