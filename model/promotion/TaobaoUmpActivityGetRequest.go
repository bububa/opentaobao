package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询营销活动 APIRequest
taobao.ump.activity.get

查询营销活动
*/
type TaobaoUmpActivityGetRequest struct {
    model.Params

    // 活动id
    actId   int64 

}

func NewTaobaoUmpActivityGetRequest() *TaobaoUmpActivityGetRequest{
    return &TaobaoUmpActivityGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpActivityGetRequest) GetApiMethodName() string {
    return "taobao.ump.activity.get"
}

func (r TaobaoUmpActivityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpActivityGetRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

func (r TaobaoUmpActivityGetRequest) GetActId() int64 {
    return r.actId
}

