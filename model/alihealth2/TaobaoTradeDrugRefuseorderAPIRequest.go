package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeDrugRefuseorderAPIRequest
阿里健康020拒单 API请求
taobao.trade.drug.refuseorder

阿里健康020拒单 */
type TaobaoTradeDrugRefuseorderAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
	// 拒单原因
	_refuseReason string
}

// New
