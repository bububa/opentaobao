package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据条件删除订阅关系 APIRequest
taobao.baichuan.items.unsubscribe.by.condition

根据条件删除订阅关系
*/
type TaobaoBaichuanItemsUnsubscribeByConditionRequest struct {
    model.Params

    // 删除条件
    condition   *Condition 

}

func NewTaobaoBaichuanItemsUnsubscribeByConditionRequest() *TaobaoBaichuanItemsUnsubscribeByConditionRequest{
    return &TaobaoBaichuanItemsUnsubscribeByConditionRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanItemsUnsubscribeByConditionRequest) GetApiMethodName() string {
    return "taobao.baichuan.items.unsubscribe.by.condition"
}

func (r TaobaoBaichuanItemsUnsubscribeByConditionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanItemsUnsubscribeByConditionRequest) SetCondition(condition *Condition) error {
    r.condition = condition
    r.Set("condition", condition)
    return nil
}

func (r TaobaoBaichuanItemsUnsubscribeByConditionRequest) GetCondition() *Condition {
    return r.condition
}

