package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除营销活动 APIRequest
taobao.ump.activity.delete

删除营销活动。对应的活动详情等将会被全部删除。
*/
type TaobaoUmpActivityDeleteRequest struct {
    model.Params

    // 活动id
    actId   int64 

}

func NewTaobaoUmpActivityDeleteRequest() *TaobaoUmpActivityDeleteRequest{
    return &TaobaoUmpActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpActivityDeleteRequest) GetApiMethodName() string {
    return "taobao.ump.activity.delete"
}

func (r TaobaoUmpActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpActivityDeleteRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

func (r TaobaoUmpActivityDeleteRequest) GetActId() int64 {
    return r.actId
}

