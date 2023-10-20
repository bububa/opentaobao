package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpcampaignfindpageAPIRequest 查询计划分页列表 API请求
// taobao.universalbp.campaign.findpage
//
// 分页查询场景内的计划列表
type TaobaouniversalbpcampaignfindpageAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// campaignQueryVO
	_campaignQueryVO *CampaignQueryVo
}

// NewTaobaouniversalbpcampaignfindpageRequest 初始化TaobaouniversalbpcampaignfindpageAPIRequest对象
func NewTaobaouniversalbpcampaignfindpageRequest() *TaobaouniversalbpcampaignfindpageAPIRequest {
	return &TaobaouniversalbpcampaignfindpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpcampaignfindpageAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.campaign.findpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpcampaignfindpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpcampaignfindpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpcampaignfindpageAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpcampaignfindpageAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetCampaignQueryVO is CampaignQueryVO Setter
// campaignQueryVO
func (r *TaobaouniversalbpcampaignfindpageAPIRequest) SetCampaignQueryVO(_campaignQueryVO *CampaignQueryVo) error {
	r._campaignQueryVO = _campaignQueryVO
	r.Set("campaign_query_v_o", _campaignQueryVO)
	return nil
}

// GetCampaignQueryVO CampaignQueryVO Getter
func (r TaobaouniversalbpcampaignfindpageAPIRequest) GetCampaignQueryVO() *CampaignQueryVo {
	return r._campaignQueryVO
}
