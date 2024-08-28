package baoxian

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baoxian"
)

// AlipayBaoxianClaimReturngoodsstatusUpdate 更新理赔单退货货物状态
// alipay.baoxian.claim.returngoodsstatus.update
//
// 更新理赔单退货货物状态
func AlipayBaoxianClaimReturngoodsstatusUpdate(ctx context.Context, clt *core.SDKClient, req *baoxian.AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest, resp *baoxian.AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
