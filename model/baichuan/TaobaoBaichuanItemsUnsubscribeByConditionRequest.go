package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据条件删除订阅关系 API请求
taobao.baichuan.items.unsubscribe.by.condition

根据条件删除订阅关系
*/
type TaobaoBaichuanItemsUnsubscribeByConditionRequest struct {
    model.Params
    // 删除条件
    _condition   *Condition
}

// 初始化TaobaoBaichuanItemsUnsubscribeByConditionRequest对象
func NewTaobaoBaichuanItemsUnsubscribeByConditionRequest() *TaobaoBaichuanItemsUnsubscribeByConditionRequest{
    return &TaobaoBaichuanItemsUnsubscribeByConditionRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemsUnsubscribeByConditionRequest) GetApiMethodName() string {
    return "taobao.baichuan.items.unsubscribe.by.condition"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemsUnsubscribeByConditionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Condition Setter
// 删除条件
func (r *TaobaoBaichuanItemsUnsubscribeByConditionRequest) SetCondition(_condition *Condition) error {
    r._condition = _condition
    r.Set("condition", _condition)
    return nil
}

// Condition Getter
func (r TaobaoBaichuanItemsUnsubscribeByConditionRequest) GetCondition() *Condition {
    return r._condition
}
