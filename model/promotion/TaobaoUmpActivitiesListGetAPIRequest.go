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
type TaobaoUmpActivitiesListGetAPIRequest struct {
    model.Params
    // 营销活动id列表。
    _ids   []int64
}

// 初始化TaobaoUmpActivitiesListGetAPIRequest对象
func NewTaobaoUmpActivitiesListGetRequest() *TaobaoUmpActivitiesListGetAPIRequest{
    return &TaobaoUmpActivitiesListGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivitiesListGetAPIRequest) GetApiMethodName() string {
    return "taobao.ump.activities.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivitiesListGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ids Setter
// 营销活动id列表。
func (r *TaobaoUmpActivitiesListGetAPIRequest) SetIds(_ids []int64) error {
    r._ids = _ids
    r.Set("ids", _ids)
    return nil
}

// Ids Getter
func (r TaobaoUmpActivitiesListGetAPIRequest) GetIds() []int64 {
    return r._ids
}
