package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除营销活动 API请求
taobao.ump.activity.delete

删除营销活动。对应的活动详情等将会被全部删除。
*/
type TaobaoUmpActivityDeleteRequest struct {
    model.Params
    // 活动id
    _actId   int64
}

// 初始化TaobaoUmpActivityDeleteRequest对象
func NewTaobaoUmpActivityDeleteRequest() *TaobaoUmpActivityDeleteRequest{
    return &TaobaoUmpActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivityDeleteRequest) GetApiMethodName() string {
    return "taobao.ump.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActId Setter
// 活动id
func (r *TaobaoUmpActivityDeleteRequest) SetActId(_actId int64) error {
    r._actId = _actId
    r.Set("act_id", _actId)
    return nil
}

// ActId Getter
func (r TaobaoUmpActivityDeleteRequest) GetActId() int64 {
    return r._actId
}
