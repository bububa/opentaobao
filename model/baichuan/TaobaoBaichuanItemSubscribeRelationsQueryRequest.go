package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按条件查询订阅关系 API请求
taobao.baichuan.item.subscribe.relations.query

按条件查询订阅关系
*/
type TaobaoBaichuanItemSubscribeRelationsQueryRequest struct {
    model.Params
    // 查询条件
    _condition   *Condition
}

// 初始化TaobaoBaichuanItemSubscribeRelationsQueryRequest对象
func NewTaobaoBaichuanItemSubscribeRelationsQueryRequest() *TaobaoBaichuanItemSubscribeRelationsQueryRequest{
    return &TaobaoBaichuanItemSubscribeRelationsQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemSubscribeRelationsQueryRequest) GetApiMethodName() string {
    return "taobao.baichuan.item.subscribe.relations.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemSubscribeRelationsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Condition Setter
// 查询条件
func (r *TaobaoBaichuanItemSubscribeRelationsQueryRequest) SetCondition(_condition *Condition) error {
    r._condition = _condition
    r.Set("condition", _condition)
    return nil
}

// Condition Getter
func (r TaobaoBaichuanItemSubscribeRelationsQueryRequest) GetCondition() *Condition {
    return r._condition
}
