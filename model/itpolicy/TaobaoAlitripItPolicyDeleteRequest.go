package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票销售规则】单条删除 APIRequest
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

func NewTaobaoAlitripItPolicyDeleteRequest() *TaobaoAlitripItPolicyDeleteRequest{
    return &TaobaoAlitripItPolicyDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripItPolicyDeleteRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.policy.delete"
}

func (r TaobaoAlitripItPolicyDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripItPolicyDeleteRequest) SetExtendAttributes(extendAttributes string) error {
    r.extendAttributes = extendAttributes
    r.Set("extend_attributes", extendAttributes)
    return nil
}

func (r TaobaoAlitripItPolicyDeleteRequest) GetExtendAttributes() string {
    return r.extendAttributes
}

func (r *TaobaoAlitripItPolicyDeleteRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

func (r TaobaoAlitripItPolicyDeleteRequest) GetOutId() string {
    return r.outId
}

func (r *TaobaoAlitripItPolicyDeleteRequest) SetTaobaoId(taobaoId int64) error {
    r.taobaoId = taobaoId
    r.Set("taobao_id", taobaoId)
    return nil
}

func (r TaobaoAlitripItPolicyDeleteRequest) GetTaobaoId() int64 {
    return r.taobaoId
}

