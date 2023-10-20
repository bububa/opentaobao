package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignFindRealCostAPIRequest 批量查询计划消耗数据 API请求
// alibaba.scbp.ad.campaign.find.real.cost
//
// 批量查询计划消耗数据
type AlibabaScbpAdCampaignFindRealCostAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 系统自动生成
	_campaignQuery *CampaignQueryDto
}

// NewAlibabaScbpAdCampaignFindRealCostRequest 初始化AlibabaScbpAdCampaignFindRealCostAPIRequest对象
func NewAlibabaScbpAdCampaignFindRealCostRequest() *AlibabaScbpAdCampaignFindRealCostAPIRequest {
	return &AlibabaScbpAdCampaignFindRealCostAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdCampaignFindRealCostAPIRequest) Reset() {
	r._topContext = nil
	r._campaignQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignFindRealCostAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.find.real.cost"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignFindRealCostAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdCampaignFindRealCostAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignFindRealCostAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdCampaignFindRealCostAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignQuery is CampaignQuery Setter
// 系统自动生成
func (r *AlibabaScbpAdCampaignFindRealCostAPIRequest) SetCampaignQuery(_campaignQuery *CampaignQueryDto) error {
	r._campaignQuery = _campaignQuery
	r.Set("campaign_query", _campaignQuery)
	return nil
}

// GetCampaignQuery CampaignQuery Getter
func (r AlibabaScbpAdCampaignFindRealCostAPIRequest) GetCampaignQuery() *CampaignQueryDto {
	return r._campaignQuery
}

var poolAlibabaScbpAdCampaignFindRealCostAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdCampaignFindRealCostRequest()
	},
}

// GetAlibabaScbpAdCampaignFindRealCostRequest 从 sync.Pool 获取 AlibabaScbpAdCampaignFindRealCostAPIRequest
func GetAlibabaScbpAdCampaignFindRealCostAPIRequest() *AlibabaScbpAdCampaignFindRealCostAPIRequest {
	return poolAlibabaScbpAdCampaignFindRealCostAPIRequest.Get().(*AlibabaScbpAdCampaignFindRealCostAPIRequest)
}

// ReleaseAlibabaScbpAdCampaignFindRealCostAPIRequest 将 AlibabaScbpAdCampaignFindRealCostAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdCampaignFindRealCostAPIRequest(v *AlibabaScbpAdCampaignFindRealCostAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdCampaignFindRealCostAPIRequest.Put(v)
}
