package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest 标签增删改 API请求
// alibaba.scbp.ad.target.tag.merge.campaign.target.tag
//
// 标签增删改
type AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest struct {
	model.Params
	// 标签数据，json格式。 最外层key：人群标签crowd/地域标签region、priceMode，第二层key: 增add、删del、改mod，第三层key：optionValue、bidRate、tagId  eg: 删除：{"crowd":{"del":[{"tagId":3595769030}]}}   修改：{"crowd":{"mod":[{"optionValue":"high_potential_order_user","bidRate":"151"}]}} 增加：{"crowd":{"add":[{"optionValue":"user_area_CA","bidRate":"133"}]}}
	_data string
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
}

// NewAlibabaScbpAdTargetTagMergeCampaignTargetTagRequest 初始化AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest对象
func NewAlibabaScbpAdTargetTagMergeCampaignTargetTagRequest() *AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest {
	return &AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.target.tag.merge.campaign.target.tag"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetData is Data Setter
// 标签数据，json格式。 最外层key：人群标签crowd/地域标签region、priceMode，第二层key: 增add、删del、改mod，第三层key：optionValue、bidRate、tagId  eg: 删除：{"crowd":{"del":[{"tagId":3595769030}]}}   修改：{"crowd":{"mod":[{"optionValue":"high_potential_order_user","bidRate":"151"}]}} 增加：{"crowd":{"add":[{"optionValue":"user_area_CA","bidRate":"133"}]}}
func (r *AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) GetData() string {
	return r._data
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
