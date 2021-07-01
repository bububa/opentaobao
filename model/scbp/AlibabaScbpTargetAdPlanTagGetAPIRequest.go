package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdPlanTagGetAPIRequest
定向推广-获取计划的定向溢价数据 API请求
alibaba.scbp.target.ad.plan.tag.get

定向推广-获取计划的定向溢价数据 */
type AlibabaScbpTargetAdPlanTagGetAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
}

// NewAlibabaScbpTargetAdPlanTagGetRequest 初始化AlibabaScbpTargetAdPlanTagGetAPIRequest对象
func NewAlibabaScbpTargetAdPlanTagGetRequest() *AlibabaScbpTargetAdPlanTagGetAPIRequest {
	return &AlibabaScbpTargetAdPlanTagGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanTagGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.tag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanTagGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampaignId Setter
// 推广计划Id
func (r *AlibabaScbpTargetAdPlanTagGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r AlibabaScbpTargetAdPlanTagGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
