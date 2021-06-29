package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标签增删改 API请求
alibaba.scbp.ad.target.tag.merge.campaign.target.tag

标签增删改
*/
type AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 标签数据，json格式。 最外层key：人群标签crowd/地域标签region、priceMode，第二层key: 增add、删del、改mod，第三层key：optionValue、bidRate、tagId  eg: 删除：{"crowd":{"del":[{"tagId":3595769030}]}}   修改：{"crowd":{"mod":[{"optionValue":"high_potential_order_user","bidRate":"151"}]}} 增加：{"crowd":{"add":[{"optionValue":"user_area_CA","bidRate":"133"}]}}
    _data   string
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest对象
func NewAlibabaScbpAdTargetTagMergeCampaignTargetTagRequest() *AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest{
    return &AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.target.tag.merge.campaign.target.tag"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) GetCampaignId() int64 {
    return r._campaignId
}
// Data Setter
// 标签数据，json格式。 最外层key：人群标签crowd/地域标签region、priceMode，第二层key: 增add、删del、改mod，第三层key：optionValue、bidRate、tagId  eg: 删除：{"crowd":{"del":[{"tagId":3595769030}]}}   修改：{"crowd":{"mod":[{"optionValue":"high_potential_order_user","bidRate":"151"}]}} 增加：{"crowd":{"add":[{"optionValue":"user_area_CA","bidRate":"133"}]}}
func (r *AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) SetData(_data string) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) GetData() string {
    return r._data
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
