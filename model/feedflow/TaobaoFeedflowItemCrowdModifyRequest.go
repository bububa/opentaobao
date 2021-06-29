package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
覆盖单元下同类型定向人群 APIRequest
taobao.feedflow.item.crowd.modify

覆盖单元下同类型定向人群
*/
type TaobaoFeedflowItemCrowdModifyRequest struct {
    model.Params

    // 人群信息
    crowds   []CrowdDto 

    // 单元id
    adgroupId   int64 

}

func NewTaobaoFeedflowItemCrowdModifyRequest() *TaobaoFeedflowItemCrowdModifyRequest{
    return &TaobaoFeedflowItemCrowdModifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCrowdModifyRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.modify"
}

func (r TaobaoFeedflowItemCrowdModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCrowdModifyRequest) SetCrowds(crowds []CrowdDto) error {
    r.crowds = crowds
    r.Set("crowds", crowds)
    return nil
}

func (r TaobaoFeedflowItemCrowdModifyRequest) GetCrowds() []CrowdDto {
    return r.crowds
}

func (r *TaobaoFeedflowItemCrowdModifyRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoFeedflowItemCrowdModifyRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

