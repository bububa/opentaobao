package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeCreateAPIRequest
创建订单 API请求
taobao.openmall.trade.create

创建Openmall订单 */
type TaobaoOpenmallTradeCreateAPIRequest struct {
	model.Params
	// 请求入参
	_paramTopTradeCreateDO *TopTradeCreateDo
}

// New
