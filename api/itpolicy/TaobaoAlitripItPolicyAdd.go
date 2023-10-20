package itpolicy

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// TaobaoAlitripItPolicyAdd 【国际机票销售规则】单条新增
// taobao.alitrip.it.policy.add
//
// 销售规则新增，成功返回taobaoId
func TaobaoAlitripItPolicyAdd(clt *core.SDKClient, req *itpolicy.TaobaoAlitripItPolicyAddAPIRequest, resp *itpolicy.TaobaoAlitripItPolicyAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
