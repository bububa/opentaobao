package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanTagGetAPIRequest 定向推广-获取计划的定向溢价数据 API请求
// alibaba.scbp.target.ad.plan.tag.get
//
// 定向推广-获取计划的定向溢价数据
type AlibabaScbpTargetAdPlanTagGetAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
}

// NewAlibabaScbpTargetAdPlanTagGetRequest 初始化AlibabaScbpTargetAdPlanTagGetAPIRequest对象
func NewAlibabaScbpTargetAdPlanTagGetRequest() *AlibabaScbpTargetAdPlanTagGetAPIRequest {
	return &AlibabaScbpTargetAdPlanTagGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpTargetAdPlanTagGetAPIRequest) Reset() {
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanTagGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.tag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanTagGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTargetAdPlanTagGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *AlibabaScbpTargetAdPlanTagGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpTargetAdPlanTagGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

var poolAlibabaScbpTargetAdPlanTagGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpTargetAdPlanTagGetRequest()
	},
}

// GetAlibabaScbpTargetAdPlanTagGetRequest 从 sync.Pool 获取 AlibabaScbpTargetAdPlanTagGetAPIRequest
func GetAlibabaScbpTargetAdPlanTagGetAPIRequest() *AlibabaScbpTargetAdPlanTagGetAPIRequest {
	return poolAlibabaScbpTargetAdPlanTagGetAPIRequest.Get().(*AlibabaScbpTargetAdPlanTagGetAPIRequest)
}

// ReleaseAlibabaScbpTargetAdPlanTagGetAPIRequest 将 AlibabaScbpTargetAdPlanTagGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanTagGetAPIRequest(v *AlibabaScbpTargetAdPlanTagGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanTagGetAPIRequest.Put(v)
}
