package scbp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) Reset() {
	r._data = ""
	r._topContext = nil
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.target.tag.merge.campaign.target.tag"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetData is Data Setter
// 标签数据，json格式。 最外层key：人群标签crowd/地域标签region、priceMode，第二层key: 增add、删del、改mod，第三层key：optionValue、bidRate、tagId  eg: 删除：{&#34;crowd&#34;:{&#34;del&#34;:[{&#34;tagId&#34;:3595769030}]}}   修改：{&#34;crowd&#34;:{&#34;mod&#34;:[{&#34;optionValue&#34;:&#34;high_potential_order_user&#34;,&#34;bidRate&#34;:&#34;151&#34;}]}} 增加：{&#34;crowd&#34;:{&#34;add&#34;:[{&#34;optionValue&#34;:&#34;user_area_CA&#34;,&#34;bidRate&#34;:&#34;133&#34;}]}}
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

var poolAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdTargetTagMergeCampaignTargetTagRequest()
	},
}

// GetAlibabaScbpAdTargetTagMergeCampaignTargetTagRequest 从 sync.Pool 获取 AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest
func GetAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest() *AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest {
	return poolAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest.Get().(*AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest)
}

// ReleaseAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest 将 AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest(v *AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest.Put(v)
}
