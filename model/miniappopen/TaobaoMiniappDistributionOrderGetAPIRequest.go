package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappDistributionOrderGetAPIRequest
小程序投放-查询小程序投放计划信息 API请求
taobao.miniapp.distribution.order.get

服务商可通过该API，读取自己开发的小程序对应的投放计划的相关信息 */
type TaobaoMiniappDistributionOrderGetAPIRequest struct {
	model.Params
	// 查询入参
	_orderIdRequest *DistributionOrderQueryByIdOpenRequest
}

// New
