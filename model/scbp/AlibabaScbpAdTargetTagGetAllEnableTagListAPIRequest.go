package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest 查询所有可添加标签信息 API请求
// alibaba.scbp.ad.target.tag.get.all.enable.tag.list
//
// 查询标签数据
type AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest struct {
	model.Params
	// 全部类型（all）,人群（crowd），地域（region）
	_type string
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
}

// NewAlibabaScbpAdTargetTagGetAllEnableTagListRequest 初始化AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest对象
func NewAlibabaScbpAdTargetTagGetAllEnableTagListRequest() *AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest {
	return &AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.target.tag.get.all.enable.tag.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 全部类型（all）,人群（crowd），地域（region）
func (r *AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest) GetType() string {
	return r._type
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdTargetTagGetAllEnableTagListAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
