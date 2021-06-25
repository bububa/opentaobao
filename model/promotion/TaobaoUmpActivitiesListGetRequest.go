package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
营销活动列表查询 APIRequest
taobao.ump.activities.list.get

按照营销活动id的列表ids，查询对应的营销活动列表。
*/
type TaobaoUmpActivitiesListGetRequest struct {
    model.Params

    // 营销活动id列表。
    ids   []Number 

}

func NewTaobaoUmpActivitiesListGetRequest() *TaobaoUmpActivitiesListGetRequest{
    return &TaobaoUmpActivitiesListGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpActivitiesListGetRequest) GetApiMethodName() string {
    return "taobao.ump.activities.list.get"
}

func (r TaobaoUmpActivitiesListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpActivitiesListGetRequest) SetIds(ids []Number) error {
    r.ids = ids
    r.Set("ids", ids)
    return nil
}

func (r TaobaoUmpActivitiesListGetRequest) GetIds() []Number {
    return r.ids
}

