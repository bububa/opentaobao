package itpolicy

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// TaobaoAlitripItPolicyAdd 【国际机票销售规则】单条新增
// taobao.alitrip.it.policy.add
//
// 销售规则新增，成功返回taobaoId
func TaobaoAlitripItPolicyAdd(clt *core.SDKClient, req *itpolicy.TaobaoAlitripItPolicyAddAPIRequest, session string) (*itpolicy.TaobaoAlitripItPolicyAddAPIResponse, error) {
	var resp itpolicy.TaobaoAlitripItPolicyAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
