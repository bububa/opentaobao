package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupFindAdGroupAPIRequest 查询推广组 API请求
// alibaba.scbp.ad.group.find.ad.group
//
// 查询推广组
type AlibabaScbpAdGroupFindAdGroupAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 入参
	_adGroupQuery *AdGroupQueryDto
}

// NewAlibabaScbpAdGroupFindAdGroupRequest 初始化AlibabaScbpAdGroupFindAdGroupAPIRequest对象
func NewAlibabaScbpAdGroupFindAdGroupRequest() *AlibabaScbpAdGroupFindAdGroupAPIRequest {
	return &AlibabaScbpAdGroupFindAdGroupAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdGroupFindAdGroupAPIRequest) Reset() {
	r._topContext = nil
	r._campaignId = 0
	r._adGroupQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupFindAdGroupAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.find.ad.group"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupFindAdGroupAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdGroupFindAdGroupAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupFindAdGroupAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdGroupFindAdGroupAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupFindAdGroupAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdGroupFindAdGroupAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdGroupQuery is AdGroupQuery Setter
// 入参
func (r *AlibabaScbpAdGroupFindAdGroupAPIRequest) SetAdGroupQuery(_adGroupQuery *AdGroupQueryDto) error {
	r._adGroupQuery = _adGroupQuery
	r.Set("ad_group_query", _adGroupQuery)
	return nil
}

// GetAdGroupQuery AdGroupQuery Getter
func (r AlibabaScbpAdGroupFindAdGroupAPIRequest) GetAdGroupQuery() *AdGroupQueryDto {
	return r._adGroupQuery
}

var poolAlibabaScbpAdGroupFindAdGroupAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdGroupFindAdGroupRequest()
	},
}

// GetAlibabaScbpAdGroupFindAdGroupRequest 从 sync.Pool 获取 AlibabaScbpAdGroupFindAdGroupAPIRequest
func GetAlibabaScbpAdGroupFindAdGroupAPIRequest() *AlibabaScbpAdGroupFindAdGroupAPIRequest {
	return poolAlibabaScbpAdGroupFindAdGroupAPIRequest.Get().(*AlibabaScbpAdGroupFindAdGroupAPIRequest)
}

// ReleaseAlibabaScbpAdGroupFindAdGroupAPIRequest 将 AlibabaScbpAdGroupFindAdGroupAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdGroupFindAdGroupAPIRequest(v *AlibabaScbpAdGroupFindAdGroupAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdGroupFindAdGroupAPIRequest.Put(v)
}
