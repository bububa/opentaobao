package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票销售规则】单条新增 APIRequest
taobao.alitrip.it.policy.add

销售规则新增，成功返回taobaoId
*/
type TaobaoAlitripItPolicyAddRequest struct {
    model.Params

    // 扩展字段
    extendAttributes   string 

    // 国际机票销售规则
    topPolicyDo   *TopPolicyDo 

}

func NewTaobaoAlitripItPolicyAddRequest() *TaobaoAlitripItPolicyAddRequest{
    return &TaobaoAlitripItPolicyAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripItPolicyAddRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.policy.add"
}

func (r TaobaoAlitripItPolicyAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripItPolicyAddRequest) SetExtendAttributes(extendAttributes string) error {
    r.extendAttributes = extendAttributes
    r.Set("extend_attributes", extendAttributes)
    return nil
}

func (r TaobaoAlitripItPolicyAddRequest) GetExtendAttributes() string {
    return r.extendAttributes
}

func (r *TaobaoAlitripItPolicyAddRequest) SetTopPolicyDo(topPolicyDo *TopPolicyDo) error {
    r.topPolicyDo = topPolicyDo
    r.Set("top_policy_do", topPolicyDo)
    return nil
}

func (r TaobaoAlitripItPolicyAddRequest) GetTopPolicyDo() *TopPolicyDo {
    return r.topPolicyDo
}

