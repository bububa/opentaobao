package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询单品单元下人群列表 APIRequest
taobao.feedflow.item.crowd.page

分页查询单品单元下人群列表
*/
type TaobaoFeedflowItemCrowdPageRequest struct {
    model.Params

    // 查询条件
    crowdQuery   *CrowdQueryDto 

}

func NewTaobaoFeedflowItemCrowdPageRequest() *TaobaoFeedflowItemCrowdPageRequest{
    return &TaobaoFeedflowItemCrowdPageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCrowdPageRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.crowd.page"
}

func (r TaobaoFeedflowItemCrowdPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCrowdPageRequest) SetCrowdQuery(crowdQuery *CrowdQueryDto) error {
    r.crowdQuery = crowdQuery
    r.Set("crowd_query", crowdQuery)
    return nil
}

func (r TaobaoFeedflowItemCrowdPageRequest) GetCrowdQuery() *CrowdQueryDto {
    return r.crowdQuery
}

