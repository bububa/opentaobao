package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCampaigngroupFindlistAPIRequest 查询计划组列表 API请求
// taobao.universalbp.campaigngroup.findlist
//
// 查询某个场景内的计划组列表
type TaobaoUniversalbpCampaigngroupFindlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// campaignGroupQueryVO
	_campaignGroupQueryVO *CampaignGroupQueryVo
}

// NewTaobaoUniversalbpCampaigngroupFindlistRequest 初始化TaobaoUniversalbpCampaigngroupFindlistAPIRequest对象
func NewTaobaoUniversalbpCampaigngroupFindlistRequest() *TaobaoUniversalbpCampaigngroupFindlistAPIRequest {
	return &TaobaoUniversalbpCampaigngroupFindlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpCampaigngroupFindlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.campaigngroup.findlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpCampaigngroupFindlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpCampaigngroupFindlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpCampaigngroupFindlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpCampaigngroupFindlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetCampaignGroupQueryVO is CampaignGroupQueryVO Setter
// campaignGroupQueryVO
func (r *TaobaoUniversalbpCampaigngroupFindlistAPIRequest) SetCampaignGroupQueryVO(_campaignGroupQueryVO *CampaignGroupQueryVo) error {
	r._campaignGroupQueryVO = _campaignGroupQueryVO
	r.Set("campaign_group_query_v_o", _campaignGroupQueryVO)
	return nil
}

// GetCampaignGroupQueryVO CampaignGroupQueryVO Getter
func (r TaobaoUniversalbpCampaigngroupFindlistAPIRequest) GetCampaignGroupQueryVO() *CampaignGroupQueryVo {
	return r._campaignGroupQueryVO
}
