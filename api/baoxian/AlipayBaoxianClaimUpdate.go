package baoxian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baoxian"
)

// AlipayBaoxianClaimUpdate 更新赔案
// alipay.baoxian.claim.update
//
// 更新保险理赔单
func AlipayBaoxianClaimUpdate(clt *core.SDKClient, req *baoxian.AlipayBaoxianClaimUpdateAPIRequest, resp *baoxian.AlipayBaoxianClaimUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
