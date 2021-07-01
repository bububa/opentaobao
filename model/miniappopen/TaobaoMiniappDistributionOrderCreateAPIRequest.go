package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappDistributionOrderCreateAPIRequest
创建小程序投放计划 API请求
taobao.miniapp.distribution.order.create

帮助商家，创建小程序的投放计划。 */
type TaobaoMiniappDistributionOrderCreateAPIRequest struct {
	model.Params
	// 投放计划信息
	_orderRequest *DistributionOrderSaveOpenRequest
}

// New
