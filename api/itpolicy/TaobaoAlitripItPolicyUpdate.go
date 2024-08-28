package itpolicy

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// TaobaoAlitripItPolicyUpdate 【国际机票销售规则】单条更新
// taobao.alitrip.it.policy.update
//
// 销售规则更新接口，可以根据taobaoId或outId修改，outId不唯一时，不能用outId修改。
func TaobaoAlitripItPolicyUpdate(ctx context.Context, clt *core.SDKClient, req *itpolicy.TaobaoAlitripItPolicyUpdateAPIRequest, resp *itpolicy.TaobaoAlitripItPolicyUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
