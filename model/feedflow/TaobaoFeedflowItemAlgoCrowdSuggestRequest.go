package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单品人群建议出价 APIRequest
taobao.feedflow.item.algo.crowd.suggest

给超级推荐的广告主查看建议出价
*/
type TaobaoFeedflowItemAlgoCrowdSuggestRequest struct {
    model.Params

    // 人群列表
    crowds   []CrowdDto 

    // 预估的宝贝id
    itemId   int64 

    // 预估的计划id
    campaignId   int64 

}

func NewTaobaoFeedflowItemAlgoCrowdSuggestRequest() *TaobaoFeedflowItemAlgoCrowdSuggestRequest{
    return &TaobaoFeedflowItemAlgoCrowdSuggestRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAlgoCrowdSuggestRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.algo.crowd.suggest"
}

func (r TaobaoFeedflowItemAlgoCrowdSuggestRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAlgoCrowdSuggestRequest) SetCrowds(crowds []CrowdDto) error {
    r.crowds = crowds
    r.Set("crowds", crowds)
    return nil
}

func (r TaobaoFeedflowItemAlgoCrowdSuggestRequest) GetCrowds() []CrowdDto {
    return r.crowds
}

func (r *TaobaoFeedflowItemAlgoCrowdSuggestRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoFeedflowItemAlgoCrowdSuggestRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoFeedflowItemAlgoCrowdSuggestRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoFeedflowItemAlgoCrowdSuggestRequest) GetCampaignId() int64 {
    return r.campaignId
}

