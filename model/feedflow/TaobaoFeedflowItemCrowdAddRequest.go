package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单品单元下，新增定向人群 APIRequest
taobao.feedflow.item.crowd.add

单品单元下，新增定向人群
*/
type TaobaoFeedflowItemCrowdAddRequest struct {
    model.Params

    // 人群列表
    crowds   []CrowdDto 

    // 单元id
    adgroupId   int64 

}

func NewTaobaoFeedflowItemCrowdAddRequest() *TaobaoFeedflowItemCrowdAddRequest{
    return &TaobaoFeedflowItemCrowdAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCrowdAddRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.add"
}

func (r TaobaoFeedflowItemCrowdAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCrowdAddRequest) SetCrowds(crowds []CrowdDto) error {
    r.crowds = crowds
    r.Set("crowds", crowds)
    return nil
}

func (r TaobaoFeedflowItemCrowdAddRequest) GetCrowds() []CrowdDto {
    return r.crowds
}

func (r *TaobaoFeedflowItemCrowdAddRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoFeedflowItemCrowdAddRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

