package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询计划列表 API请求
taobao.feedflow.item.campaign.page

批量查询计划列表
*/
type TaobaoFeedflowItemCampaignPageRequest struct {
    model.Params
    // 入参
    _campaignQuery   *CampaignQueryDTO
}

// 初始化TaobaoFeedflowItemCampaignPageRequest对象
func NewTaobaoFeedflowItemCampaignPageRequest() *TaobaoFeedflowItemCampaignPageRequest{
    return &TaobaoFeedflowItemCampaignPageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignPageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignQuery Setter
// 入参
func (r *TaobaoFeedflowItemCampaignPageRequest) SetCampaignQuery(_campaignQuery *CampaignQueryDTO) error {
    r._campaignQuery = _campaignQuery
    r.Set("campaign_query", _campaignQuery)
    return nil
}

// CampaignQuery Getter
func (r TaobaoFeedflowItemCampaignPageRequest) GetCampaignQuery() *CampaignQueryDTO {
    return r._campaignQuery
}
