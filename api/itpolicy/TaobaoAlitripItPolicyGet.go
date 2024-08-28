package itpolicy

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// TaobaoAlitripItPolicyGet 【国际机票销售规则】单条查询
// taobao.alitrip.it.policy.get
//
// 通过此接口可以查询单条销售规则的详情，可以根据taobaoId或outId查询，用户outId查询时，如果outId不唯一，只返回最新添加的一条数据。taobaoId为新增成功时候返回的唯一id，outId为新增时的policy_id（产品编号）
func TaobaoAlitripItPolicyGet(ctx context.Context, clt *core.SDKClient, req *itpolicy.TaobaoAlitripItPolicyGetAPIRequest, resp *itpolicy.TaobaoAlitripItPolicyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
