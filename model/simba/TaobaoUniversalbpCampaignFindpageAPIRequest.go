package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCampaignFindpageAPIRequest 查询计划分页列表 API请求
// taobao.universalbp.campaign.findpage
//
// 分页查询场景内的计划列表
type TaobaoUniversalbpCampaignFindpageAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// campaignQueryVO
	_campaignQueryVO *CampaignQueryVo
}

// NewTaobaoUniversalbpCampaignFindpageRequest 初始化TaobaoUniversalbpCampaignFindpageAPIRequest对象
func NewTaobaoUniversalbpCampaignFindpageRequest() *TaobaoUniversalbpCampaignFindpageAPIRequest {
	return &TaobaoUniversalbpCampaignFindpageAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpCampaignFindpageAPIRequest) Reset() {
	r._topServiceContext = nil
	r._campaignQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpCampaignFindpageAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.campaign.findpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpCampaignFindpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpCampaignFindpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpCampaignFindpageAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpCampaignFindpageAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetCampaignQueryVO is CampaignQueryVO Setter
// campaignQueryVO
func (r *TaobaoUniversalbpCampaignFindpageAPIRequest) SetCampaignQueryVO(_campaignQueryVO *CampaignQueryVo) error {
	r._campaignQueryVO = _campaignQueryVO
	r.Set("campaign_query_v_o", _campaignQueryVO)
	return nil
}

// GetCampaignQueryVO CampaignQueryVO Getter
func (r TaobaoUniversalbpCampaignFindpageAPIRequest) GetCampaignQueryVO() *CampaignQueryVo {
	return r._campaignQueryVO
}

var poolTaobaoUniversalbpCampaignFindpageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpCampaignFindpageRequest()
	},
}

// GetTaobaoUniversalbpCampaignFindpageRequest 从 sync.Pool 获取 TaobaoUniversalbpCampaignFindpageAPIRequest
func GetTaobaoUniversalbpCampaignFindpageAPIRequest() *TaobaoUniversalbpCampaignFindpageAPIRequest {
	return poolTaobaoUniversalbpCampaignFindpageAPIRequest.Get().(*TaobaoUniversalbpCampaignFindpageAPIRequest)
}

// ReleaseTaobaoUniversalbpCampaignFindpageAPIRequest 将 TaobaoUniversalbpCampaignFindpageAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpCampaignFindpageAPIRequest(v *TaobaoUniversalbpCampaignFindpageAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpCampaignFindpageAPIRequest.Put(v)
}
