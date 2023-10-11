package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest 查询无界版计划对应的原场景计划id API请求
// taobao.universalbp.campaign.findsubcampaignid
//
// 查询该场景下，无界版计划对应的原场景的计划
type TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// long
	_tpLong int64
}

// NewTaobaoUniversalbpCampaignFindsubcampaignidRequest 初始化TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest对象
func NewTaobaoUniversalbpCampaignFindsubcampaignidRequest() *TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest {
	return &TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.campaign.findsubcampaignid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetTpLong is TpLong Setter
// long
func (r *TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest) SetTpLong(_tpLong int64) error {
	r._tpLong = _tpLong
	r.Set("tp_long", _tpLong)
	return nil
}

// GetTpLong TpLong Getter
func (r TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest) GetTpLong() int64 {
	return r._tpLong
}
