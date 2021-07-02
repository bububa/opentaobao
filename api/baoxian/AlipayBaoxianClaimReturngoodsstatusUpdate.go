package baoxian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baoxian"
)

// AlipayBaoxianClaimReturngoodsstatusUpdate 更新理赔单退货货物状态
// alipay.baoxian.claim.returngoodsstatus.update
//
// 更新理赔单退货货物状态
func AlipayBaoxianClaimReturngoodsstatusUpdate(clt *core.SDKClient, req *baoxian.AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest, session string) (*baoxian.AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse, error) {
	var resp baoxian.AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
