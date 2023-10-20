package baoxian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baoxian"
)

// AlipayBaoxianClaimReturngoodsstatusUpdate 更新理赔单退货货物状态
// alipay.baoxian.claim.returngoodsstatus.update
//
// 更新理赔单退货货物状态
func AlipayBaoxianClaimReturngoodsstatusUpdate(clt *core.SDKClient, req *baoxian.AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest, resp *baoxian.AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
