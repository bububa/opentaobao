package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票销售规则】单条查询 APIRequest
taobao.alitrip.it.policy.get

通过此接口可以查询单条销售规则的详情，可以根据taobaoId或outId查询，用户outId查询时，如果outId不唯一，只返回最新添加的一条数据。taobaoId为新增成功时候返回的唯一id，outId为新增时的policy_id（产品编号）
*/
type TaobaoAlitripItPolicyGetRequest struct {
    model.Params

    // 预留扩展字段
    extendAttributes   string 

    // 接入方产品编号
    outId   string 

    // 淘宝政策id
    taobaoId   int64 

}

func NewTaobaoAlitripItPolicyGetRequest() *TaobaoAlitripItPolicyGetRequest{
    return &TaobaoAlitripItPolicyGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripItPolicyGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.policy.get"
}

func (r TaobaoAlitripItPolicyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripItPolicyGetRequest) SetExtendAttributes(extendAttributes string) error {
    r.extendAttributes = extendAttributes
    r.Set("extend_attributes", extendAttributes)
    return nil
}

func (r TaobaoAlitripItPolicyGetRequest) GetExtendAttributes() string {
    return r.extendAttributes
}

func (r *TaobaoAlitripItPolicyGetRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

func (r TaobaoAlitripItPolicyGetRequest) GetOutId() string {
    return r.outId
}

func (r *TaobaoAlitripItPolicyGetRequest) SetTaobaoId(taobaoId int64) error {
    r.taobaoId = taobaoId
    r.Set("taobao_id", taobaoId)
    return nil
}

func (r TaobaoAlitripItPolicyGetRequest) GetTaobaoId() int64 {
    return r.taobaoId
}

