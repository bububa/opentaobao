package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
营销活动列表查询 API请求
taobao.ump.activities.list.get

按照营销活动id的列表ids，查询对应的营销活动列表。
*/
type TaobaoUmpActivitiesListGetRequest struct {
    model.Params
    // 营销活动id列表。
    ids   []int64
}

// 初始化TaobaoUmpActivitiesListGetRequest对象
func NewTaobaoUmpActivitiesListGetRequest() *TaobaoUmpActivitiesListGetRequest{
    return &TaobaoUmpActivitiesListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivitiesListGetRequest) GetApiMethodName() string {
    return "taobao.ump.activities.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivitiesListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ids Setter
// 营销活动id列表。
func (r *TaobaoUmpActivitiesListGetRequest) SetIds(ids []int64) error {
    r.ids = ids
    r.Set("ids", ids)
    return nil
}

// Ids Getter
func (r TaobaoUmpActivitiesListGetRequest) GetIds() []int64 {
    return r.ids
}
