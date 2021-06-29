package itpolicy

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/itpolicy"
)

/* 
【国际机票销售规则】单条删除 
taobao.alitrip.it.policy.delete

销售规则删除接口，可以根据taobaoId或outId删除，根据outId删除时，如果outId不唯一，返回失败
*/
func TaobaoAlitripItPolicyDelete(clt *core.SDKClient, req *itpolicy.TaobaoAlitripItPolicyDeleteRequest, session string) (*itpolicy.TaobaoAlitripItPolicyDeleteAPIResponse, error) {
    var resp itpolicy.TaobaoAlitripItPolicyDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
