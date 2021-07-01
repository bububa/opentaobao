package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappDistributionOrderItemsBindAPIRequest
小程序投放-基于投放计划绑定/解绑商品 API请求
taobao.miniapp.distribution.order.items.bind

提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。 */
type TaobaoMiniappDistributionOrderItemsBindAPIRequest struct {
	model.Params
	// 商品id列表
	_targetEntityList []string
	// true表示新增绑定，false表示解绑
	_addBind bool
	// 投放计划标识id
	_distributeId int64
}

// New
