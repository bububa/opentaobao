package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改人群出价或状态 APIRequest
taobao.feedflow.item.crowd.modifybind

修改人群出价或状态
*/
type TaobaoFeedflowItemCrowdModifybindRequest struct {
    model.Params

    // 人群信息
    crowds   []CrowdDto 

    // 单元id
    adgroupId   int64 

}

func NewTaobaoFeedflowItemCrowdModifybindRequest() *TaobaoFeedflowItemCrowdModifybindRequest{
    return &TaobaoFeedflowItemCrowdModifybindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCrowdModifybindRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.modifybind"
}

func (r TaobaoFeedflowItemCrowdModifybindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCrowdModifybindRequest) SetCrowds(crowds []CrowdDto) error {
    r.crowds = crowds
    r.Set("crowds", crowds)
    return nil
}

func (r TaobaoFeedflowItemCrowdModifybindRequest) GetCrowds() []CrowdDto {
    return r.crowds
}

func (r *TaobaoFeedflowItemCrowdModifybindRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoFeedflowItemCrowdModifybindRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

