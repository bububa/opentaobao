package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeCustomizationRefundEnableAPIRequest
定制订单设置允许仅退款 API请求
taobao.opentrade.customization.refund.enable

定制订单设置允许仅退款 */
type TaobaoOpentradeCustomizationRefundEnableAPIRequest struct {
	model.Params
	// 主订单ID
	_tradeId int64
}

// New
