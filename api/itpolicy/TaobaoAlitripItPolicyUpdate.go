package itpolicy

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/itpolicy"
)

/* 
【国际机票销售规则】单条更新 
taobao.alitrip.it.policy.update

销售规则更新接口，可以根据taobaoId或outId修改，outId不唯一时，不能用outId修改。
*/
func TaobaoAlitripItPolicyUpdate(clt *core.SDKClient, req *itpolicy.TaobaoAlitripItPolicyUpdateRequest, session string) (*itpolicy.TaobaoAlitripItPolicyUpdateAPIResponse, error) {
    var resp itpolicy.TaobaoAlitripItPolicyUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}