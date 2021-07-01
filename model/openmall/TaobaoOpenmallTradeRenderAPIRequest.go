package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeRenderAPIRequest
渲染订单价格 API请求
taobao.openmall.trade.render

请求渲染订单价格 */
type TaobaoOpenmallTradeRenderAPIRequest struct {
	model.Params
	// 请求入参
	_paramTopTradeCreateDO *TopTradeCreateDo
}

// New
