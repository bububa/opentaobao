package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomWttOpentradeCreateorderAPIRequest
充值送活动下单接口 API请求
alibaba.alicom.wtt.opentrade.createorder

提供给话费宝创建淘宝订单 */
type AlibabaAlicomWttOpentradeCreateorderAPIRequest struct {
	model.Params
	// 入参请求说明
	_param0 *OpentradCreateOrderRequestDto
}

// New
