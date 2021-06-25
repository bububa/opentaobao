package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按条件查询订阅关系 APIRequest
taobao.baichuan.item.subscribe.relations.query

按条件查询订阅关系
*/
type TaobaoBaichuanItemSubscribeRelationsQueryRequest struct {
    model.Params

    // 查询条件
    condition   *Condition 

}

func NewTaobaoBaichuanItemSubscribeRelationsQueryRequest() *TaobaoBaichuanItemSubscribeRelationsQueryRequest{
    return &TaobaoBaichuanItemSubscribeRelationsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanItemSubscribeRelationsQueryRequest) GetApiMethodName() string {
    return "taobao.baichuan.item.subscribe.relations.query"
}

func (r TaobaoBaichuanItemSubscribeRelationsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanItemSubscribeRelationsQueryRequest) SetCondition(condition *Condition) error {
    r.condition = condition
    r.Set("condition", condition)
    return nil
}

func (r TaobaoBaichuanItemSubscribeRelationsQueryRequest) GetCondition() *Condition {
    return r.condition
}

