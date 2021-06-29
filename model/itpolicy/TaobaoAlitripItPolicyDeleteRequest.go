package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票销售规则】单条删除 API请求
taobao.alitrip.it.policy.delete

销售规则删除接口，可以根据taobaoId或outId删除，根据outId删除时，如果outId不唯一，返回失败
*/
type TaobaoAlitripItPolicyDeleteRequest struct {
    model.Params
    // 扩展字段
    extendAttributes   string
    // 接入方产品id
    outId   string
    // 淘宝政策id
    taobaoId   int64
}

// 初始化TaobaoAlitripItPolicyDeleteRequest对象
func NewTaobaoAlitripItPolicyDeleteRequest() *TaobaoAlitripItPolicyDeleteRequest{
    return &TaobaoAlitripItPolicyDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItPolicyDeleteRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.policy.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItPolicyDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtendAttributes Setter
// 扩展字段
func (r *TaobaoAlitripItPolicyDeleteRequest) SetExtendAttributes(extendAttributes string) error {
    r.extendAttributes = extendAttributes
    r.Set("extend_attributes", extendAttributes)
    return nil
}

// ExtendAttributes Getter
func (r TaobaoAlitripItPolicyDeleteRequest) GetExtendAttributes() string {
    return r.extendAttributes
}
// OutId Setter
// 接入方产品id
func (r *TaobaoAlitripItPolicyDeleteRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

// OutId Getter
func (r TaobaoAlitripItPolicyDeleteRequest) GetOutId() string {
    return r.outId
}
// TaobaoId Setter
// 淘宝政策id
func (r *TaobaoAlitripItPolicyDeleteRequest) SetTaobaoId(taobaoId int64) error {
    r.taobaoId = taobaoId
    r.Set("taobao_id", taobaoId)
    return nil
}

// TaobaoId Getter
func (r TaobaoAlitripItPolicyDeleteRequest) GetTaobaoId() int64 {
    return r.taobaoId
}
