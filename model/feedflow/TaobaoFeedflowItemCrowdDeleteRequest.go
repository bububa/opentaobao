package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品人群 APIRequest
taobao.feedflow.item.crowd.delete

删除单品人群
*/
type TaobaoFeedflowItemCrowdDeleteRequest struct {
    model.Params

    // 人群结构
    crowds   []CrowdDto 

    // 单元id
    adgroupId   int64 

}

func NewTaobaoFeedflowItemCrowdDeleteRequest() *TaobaoFeedflowItemCrowdDeleteRequest{
    return &TaobaoFeedflowItemCrowdDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCrowdDeleteRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.delete"
}

func (r TaobaoFeedflowItemCrowdDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCrowdDeleteRequest) SetCrowds(crowds []CrowdDto) error {
    r.crowds = crowds
    r.Set("crowds", crowds)
    return nil
}

func (r TaobaoFeedflowItemCrowdDeleteRequest) GetCrowds() []CrowdDto {
    return r.crowds
}

func (r *TaobaoFeedflowItemCrowdDeleteRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoFeedflowItemCrowdDeleteRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

