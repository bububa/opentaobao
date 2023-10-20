package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignFindCampaignPageAPIRequest 分页查询计划 API请求
// alibaba.scbp.ad.campaign.find.campaign.page
//
// 分页查询计划
type AlibabaScbpAdCampaignFindCampaignPageAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求实体类
	_campaignQuery *CampaignQueryDto
}

// NewAlibabaScbpAdCampaignFindCampaignPageRequest 初始化AlibabaScbpAdCampaignFindCampaignPageAPIRequest对象
func NewAlibabaScbpAdCampaignFindCampaignPageRequest() *AlibabaScbpAdCampaignFindCampaignPageAPIRequest {
	return &AlibabaScbpAdCampaignFindCampaignPageAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdCampaignFindCampaignPageAPIRequest) Reset() {
	r._topContext = nil
	r._campaignQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignFindCampaignPageAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.find.campaign.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignFindCampaignPageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdCampaignFindCampaignPageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignFindCampaignPageAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdCampaignFindCampaignPageAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignQuery is CampaignQuery Setter
// 请求实体类
func (r *AlibabaScbpAdCampaignFindCampaignPageAPIRequest) SetCampaignQuery(_campaignQuery *CampaignQueryDto) error {
	r._campaignQuery = _campaignQuery
	r.Set("campaign_query", _campaignQuery)
	return nil
}

// GetCampaignQuery CampaignQuery Getter
func (r AlibabaScbpAdCampaignFindCampaignPageAPIRequest) GetCampaignQuery() *CampaignQueryDto {
	return r._campaignQuery
}

var poolAlibabaScbpAdCampaignFindCampaignPageAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdCampaignFindCampaignPageRequest()
	},
}

// GetAlibabaScbpAdCampaignFindCampaignPageRequest 从 sync.Pool 获取 AlibabaScbpAdCampaignFindCampaignPageAPIRequest
func GetAlibabaScbpAdCampaignFindCampaignPageAPIRequest() *AlibabaScbpAdCampaignFindCampaignPageAPIRequest {
	return poolAlibabaScbpAdCampaignFindCampaignPageAPIRequest.Get().(*AlibabaScbpAdCampaignFindCampaignPageAPIRequest)
}

// ReleaseAlibabaScbpAdCampaignFindCampaignPageAPIRequest 将 AlibabaScbpAdCampaignFindCampaignPageAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdCampaignFindCampaignPageAPIRequest(v *AlibabaScbpAdCampaignFindCampaignPageAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdCampaignFindCampaignPageAPIRequest.Put(v)
}
