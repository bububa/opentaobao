package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpcampaignfindsubcampaignidAPIRequest 查询无界版计划对应的原场景计划id API请求
// taobao.universalbp.campaign.findsubcampaignid
//
// 查询该场景下，无界版计划对应的原场景的计划
type TaobaouniversalbpcampaignfindsubcampaignidAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// long
	_tpLong int64
}

// NewTaobaouniversalbpcampaignfindsubcampaignidRequest 初始化TaobaouniversalbpcampaignfindsubcampaignidAPIRequest对象
func NewTaobaouniversalbpcampaignfindsubcampaignidRequest() *TaobaouniversalbpcampaignfindsubcampaignidAPIRequest {
	return &TaobaouniversalbpcampaignfindsubcampaignidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpcampaignfindsubcampaignidAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.campaign.findsubcampaignid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpcampaignfindsubcampaignidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpcampaignfindsubcampaignidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpcampaignfindsubcampaignidAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpcampaignfindsubcampaignidAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTpLong is TpLong Setter
// long
func (r *TaobaouniversalbpcampaignfindsubcampaignidAPIRequest) SetTpLong(_tpLong int64) error {
	r._tpLong = _tpLong
	r.Set("tp_long", _tpLong)
	return nil
}

// GetTpLong TpLong Getter
func (r TaobaouniversalbpcampaignfindsubcampaignidAPIRequest) GetTpLong() int64 {
	return r._tpLong
}
