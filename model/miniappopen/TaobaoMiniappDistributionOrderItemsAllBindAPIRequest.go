package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappDistributionOrderItemsAllBindAPIRequest
小程序投放-基于投放计划绑定/解绑全店商品 API请求
taobao.miniapp.distribution.order.items.all.bind

提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。该接口对应的能力是全店投放，该能力的使用，需要联系平台运营进行人工申请，申请通过以后方可调用成功。 */
type TaobaoMiniappDistributionOrderItemsAllBindAPIRequest struct {
	model.Params
	// 绑定/解绑的入参信息
	_allItemBindRequest *DistributionOrderBindTargetEntityOpenRequestV2
}

// New
