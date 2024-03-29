package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanListAPIRequest 定向推广-查询定向推广计划列表并返回计划基础信息 API请求
// alibaba.scbp.target.ad.plan.list
//
// 定向推广-查询定向推广计划列表并返回计划基础信息
type AlibabaScbpTargetAdPlanListAPIRequest struct {
	model.Params
	// TopP4pQuickCampaignQuery
	_topP4pQuickCampaignQuery *TopP4pQuickCampaignQueryDto
}

// NewAlibabaScbpTargetAdPlanListRequest 初始化AlibabaScbpTargetAdPlanListAPIRequest对象
func NewAlibabaScbpTargetAdPlanListRequest() *AlibabaScbpTargetAdPlanListAPIRequest {
	return &AlibabaScbpTargetAdPlanListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpTargetAdPlanListAPIRequest) Reset() {
	r._topP4pQuickCampaignQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanListAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTargetAdPlanListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopP4pQuickCampaignQuery is TopP4pQuickCampaignQuery Setter
// TopP4pQuickCampaignQuery
func (r *AlibabaScbpTargetAdPlanListAPIRequest) SetTopP4pQuickCampaignQuery(_topP4pQuickCampaignQuery *TopP4pQuickCampaignQueryDto) error {
	r._topP4pQuickCampaignQuery = _topP4pQuickCampaignQuery
	r.Set("top_p4p_quick_campaign_query", _topP4pQuickCampaignQuery)
	return nil
}

// GetTopP4pQuickCampaignQuery TopP4pQuickCampaignQuery Getter
func (r AlibabaScbpTargetAdPlanListAPIRequest) GetTopP4pQuickCampaignQuery() *TopP4pQuickCampaignQueryDto {
	return r._topP4pQuickCampaignQuery
}

var poolAlibabaScbpTargetAdPlanListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpTargetAdPlanListRequest()
	},
}

// GetAlibabaScbpTargetAdPlanListRequest 从 sync.Pool 获取 AlibabaScbpTargetAdPlanListAPIRequest
func GetAlibabaScbpTargetAdPlanListAPIRequest() *AlibabaScbpTargetAdPlanListAPIRequest {
	return poolAlibabaScbpTargetAdPlanListAPIRequest.Get().(*AlibabaScbpTargetAdPlanListAPIRequest)
}

// ReleaseAlibabaScbpTargetAdPlanListAPIRequest 将 AlibabaScbpTargetAdPlanListAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanListAPIRequest(v *AlibabaScbpTargetAdPlanListAPIRequest) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanListAPIRequest.Put(v)
}
