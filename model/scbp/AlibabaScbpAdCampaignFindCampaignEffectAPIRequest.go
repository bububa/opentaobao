package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignFindCampaignEffectAPIRequest 批量查询计划效果数据 API请求
// alibaba.scbp.ad.campaign.find.campaign.effect
//
// 批量查询计划效果数据
type AlibabaScbpAdCampaignFindCampaignEffectAPIRequest struct {
	model.Params
	// 计划id集合
	_campaignIdList []string
	// 开始时间
	_beginDate string
	// 结束时间
	_endDate string
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdCampaignFindCampaignEffectRequest 初始化AlibabaScbpAdCampaignFindCampaignEffectAPIRequest对象
func NewAlibabaScbpAdCampaignFindCampaignEffectRequest() *AlibabaScbpAdCampaignFindCampaignEffectAPIRequest {
	return &AlibabaScbpAdCampaignFindCampaignEffectAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) Reset() {
	r._campaignIdList = r._campaignIdList[:0]
	r._beginDate = ""
	r._endDate = ""
	r._topContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.find.campaign.effect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignIdList is CampaignIdList Setter
// 计划id集合
func (r *AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) SetCampaignIdList(_campaignIdList []string) error {
	r._campaignIdList = _campaignIdList
	r.Set("campaign_id_list", _campaignIdList)
	return nil
}

// GetCampaignIdList CampaignIdList Getter
func (r AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) GetCampaignIdList() []string {
	return r._campaignIdList
}

// SetBeginDate is BeginDate Setter
// 开始时间
func (r *AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

var poolAlibabaScbpAdCampaignFindCampaignEffectAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdCampaignFindCampaignEffectRequest()
	},
}

// GetAlibabaScbpAdCampaignFindCampaignEffectRequest 从 sync.Pool 获取 AlibabaScbpAdCampaignFindCampaignEffectAPIRequest
func GetAlibabaScbpAdCampaignFindCampaignEffectAPIRequest() *AlibabaScbpAdCampaignFindCampaignEffectAPIRequest {
	return poolAlibabaScbpAdCampaignFindCampaignEffectAPIRequest.Get().(*AlibabaScbpAdCampaignFindCampaignEffectAPIRequest)
}

// ReleaseAlibabaScbpAdCampaignFindCampaignEffectAPIRequest 将 AlibabaScbpAdCampaignFindCampaignEffectAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdCampaignFindCampaignEffectAPIRequest(v *AlibabaScbpAdCampaignFindCampaignEffectAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdCampaignFindCampaignEffectAPIRequest.Put(v)
}
