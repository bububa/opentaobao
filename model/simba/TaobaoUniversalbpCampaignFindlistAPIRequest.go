package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpcampaignfindlistAPIRequest 查询全量计划列表(不分页) API请求
// taobao.universalbp.campaign.findlist
//
// 查询场景内的全量计划列表
type TaobaouniversalbpcampaignfindlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// campaignQueryVO
	_campaignQueryVO *CampaignQueryVo
}

// NewTaobaouniversalbpcampaignfindlistRequest 初始化TaobaouniversalbpcampaignfindlistAPIRequest对象
func NewTaobaouniversalbpcampaignfindlistRequest() *TaobaouniversalbpcampaignfindlistAPIRequest {
	return &TaobaouniversalbpcampaignfindlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpcampaignfindlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.campaign.findlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpcampaignfindlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpcampaignfindlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpcampaignfindlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpcampaignfindlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetCampaignQueryVO is CampaignQueryVO Setter
// campaignQueryVO
func (r *TaobaouniversalbpcampaignfindlistAPIRequest) SetCampaignQueryVO(_campaignQueryVO *CampaignQueryVo) error {
	r._campaignQueryVO = _campaignQueryVO
	r.Set("campaign_query_v_o", _campaignQueryVO)
	return nil
}

// GetCampaignQueryVO CampaignQueryVO Getter
func (r TaobaouniversalbpcampaignfindlistAPIRequest) GetCampaignQueryVO() *CampaignQueryVo {
	return r._campaignQueryVO
}
