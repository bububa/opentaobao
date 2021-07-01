package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstVendingShippingCallbackAPIRequest
售货机出货回传接口 API请求
alibaba.lst.vending.shipping.callback

零售通自动售货机商品出货回传接口，同步商品出库最新状态。 */
type AlibabaLstVendingShippingCallbackAPIRequest struct {
	model.Params
	// 厂商设备编码
	_equipmentCode string
	// 交易流水号
	_tradeFlowNo string
	// 处理结果代码
	_code string
	// 处理结果代码描述
	_message string
	// 出货时间
	_time string
}

// New
