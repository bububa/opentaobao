package baoxian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest
更新理赔单退货货物状态 API请求
alipay.baoxian.claim.returngoodsstatus.update

更新理赔单退货货物状态 */
type AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest struct {
	model.Params
	// 理赔单号
	_claimNo string
	// 退货货物状态
	_goodsStatus string
}

// New
