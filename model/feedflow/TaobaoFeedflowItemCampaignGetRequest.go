package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过计划id查询计划 API请求
taobao.feedflow.item.campaign.get

通过计划id查询计划
*/
type TaobaoFeedflowItemCampaignGetAPIRequest struct {
    model.Params
    // 计划id
    _campaginId   int64
}

// 初始化TaobaoFeedflowItemCampaignGetAPIRequest对象
func NewTaobaoFeedflowItemCampaignGetRequest() *TaobaoFeedflowItemCampaignGetAPIRequest{
    return &TaobaoFeedflowItemCampaignGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignGetAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaginId Setter
// 计划id
func (r *TaobaoFeedflowItemCampaignGetAPIRequest) SetCampaginId(_campaginId int64) error {
    r._campaginId = _campaginId
    r.Set("campagin_id", _campaginId)
    return nil
}

// CampaginId Getter
func (r TaobaoFeedflowItemCampaignGetAPIRequest) GetCampaginId() int64 {
    return r._campaginId
}
