package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询屏蔽品 API请求
alibaba.scbp.ad.group.find.forbidden.product

查询屏蔽品
*/
type AlibabaScbpAdGroupFindForbiddenProductRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdGroupFindForbiddenProductRequest对象
func NewAlibabaScbpAdGroupFindForbiddenProductRequest() *AlibabaScbpAdGroupFindForbiddenProductRequest{
    return &AlibabaScbpAdGroupFindForbiddenProductRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupFindForbiddenProductRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.find.forbidden.product"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupFindForbiddenProductRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupFindForbiddenProductRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupFindForbiddenProductRequest) GetCampaignId() int64 {
    return r._campaignId
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupFindForbiddenProductRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupFindForbiddenProductRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
