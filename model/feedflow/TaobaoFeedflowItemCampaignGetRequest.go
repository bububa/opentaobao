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
type TaobaoFeedflowItemCampaignGetRequest struct {
    model.Params
    // 计划id
    _campaginId   int64
}

// 初始化TaobaoFeedflowItemCampaignGetRequest对象
func NewTaobaoFeedflowItemCampaignGetRequest() *TaobaoFeedflowItemCampaignGetRequest{
    return &TaobaoFeedflowItemCampaignGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignGetRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaginId Setter
// 计划id
func (r *TaobaoFeedflowItemCampaignGetRequest) SetCampaginId(_campaginId int64) error {
    r._campaginId = _campaginId
    r.Set("campagin_id", _campaginId)
    return nil
}

// CampaginId Getter
func (r TaobaoFeedflowItemCampaignGetRequest) GetCampaginId() int64 {
    return r._campaginId
}
