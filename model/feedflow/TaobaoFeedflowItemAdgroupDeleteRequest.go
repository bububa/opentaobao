package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据单元id删除单元 APIRequest
taobao.feedflow.item.adgroup.delete

根据单元id删除单元
*/
type TaobaoFeedflowItemAdgroupDeleteRequest struct {
    model.Params

    // 计划id
    campaignId   int64 

    // 单元id列表
    adgroupIdList   []int64 

}

func NewTaobaoFeedflowItemAdgroupDeleteRequest() *TaobaoFeedflowItemAdgroupDeleteRequest{
    return &TaobaoFeedflowItemAdgroupDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdgroupDeleteRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.delete"
}

func (r TaobaoFeedflowItemAdgroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdgroupDeleteRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoFeedflowItemAdgroupDeleteRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoFeedflowItemAdgroupDeleteRequest) SetAdgroupIdList(adgroupIdList []int64) error {
    r.adgroupIdList = adgroupIdList
    r.Set("adgroup_id_list", adgroupIdList)
    return nil
}

func (r TaobaoFeedflowItemAdgroupDeleteRequest) GetAdgroupIdList() []int64 {
    return r.adgroupIdList
}

