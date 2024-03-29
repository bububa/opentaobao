package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdCampaignEffectAPIRequest 定向推广-获取计划维度推广效果 API请求
// alibaba.scbp.target.ad.campaign.effect
//
// 定向推广-获取计划维度推广效果
type AlibabaScbpTargetAdCampaignEffectAPIRequest struct {
	model.Params
	// 结束时间 当inteval=7或30的时候 不需要填写，当inteval=1时需要填写（开始结束时间区间不允许大于180天）
	_endDate string
	// 开始时间 当inteval=7或30的时候 不需要填写，当inteval=1时需要填写（开始结束时间区间不允许大于180天）
	_beginDate string
	// 统计区间 只能为1 7 30
	_interval int64
	// 当填写时，展示指定id的数据，不填写，则展示全部计划总数据
	_campaignId int64
}

// NewAlibabaScbpTargetAdCampaignEffectRequest 初始化AlibabaScbpTargetAdCampaignEffectAPIRequest对象
func NewAlibabaScbpTargetAdCampaignEffectRequest() *AlibabaScbpTargetAdCampaignEffectAPIRequest {
	return &AlibabaScbpTargetAdCampaignEffectAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpTargetAdCampaignEffectAPIRequest) Reset() {
	r._endDate = ""
	r._beginDate = ""
	r._interval = 0
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdCampaignEffectAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.campaign.effect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdCampaignEffectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTargetAdCampaignEffectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndDate is EndDate Setter
// 结束时间 当inteval=7或30的时候 不需要填写，当inteval=1时需要填写（开始结束时间区间不允许大于180天）
func (r *AlibabaScbpTargetAdCampaignEffectAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaScbpTargetAdCampaignEffectAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetBeginDate is BeginDate Setter
// 开始时间 当inteval=7或30的时候 不需要填写，当inteval=1时需要填写（开始结束时间区间不允许大于180天）
func (r *AlibabaScbpTargetAdCampaignEffectAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaScbpTargetAdCampaignEffectAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetInterval is Interval Setter
// 统计区间 只能为1 7 30
func (r *AlibabaScbpTargetAdCampaignEffectAPIRequest) SetInterval(_interval int64) error {
	r._interval = _interval
	r.Set("interval", _interval)
	return nil
}

// GetInterval Interval Getter
func (r AlibabaScbpTargetAdCampaignEffectAPIRequest) GetInterval() int64 {
	return r._interval
}

// SetCampaignId is CampaignId Setter
// 当填写时，展示指定id的数据，不填写，则展示全部计划总数据
func (r *AlibabaScbpTargetAdCampaignEffectAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpTargetAdCampaignEffectAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

var poolAlibabaScbpTargetAdCampaignEffectAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpTargetAdCampaignEffectRequest()
	},
}

// GetAlibabaScbpTargetAdCampaignEffectRequest 从 sync.Pool 获取 AlibabaScbpTargetAdCampaignEffectAPIRequest
func GetAlibabaScbpTargetAdCampaignEffectAPIRequest() *AlibabaScbpTargetAdCampaignEffectAPIRequest {
	return poolAlibabaScbpTargetAdCampaignEffectAPIRequest.Get().(*AlibabaScbpTargetAdCampaignEffectAPIRequest)
}

// ReleaseAlibabaScbpTargetAdCampaignEffectAPIRequest 将 AlibabaScbpTargetAdCampaignEffectAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpTargetAdCampaignEffectAPIRequest(v *AlibabaScbpTargetAdCampaignEffectAPIRequest) {
	v.Reset()
	poolAlibabaScbpTargetAdCampaignEffectAPIRequest.Put(v)
}
