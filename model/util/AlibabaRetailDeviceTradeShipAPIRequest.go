package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailDeviceTradeShipAPIRequest
贩卖机掉货成功通知 API请求
alibaba.retail.device.trade.ship

贩卖机发货 */
type AlibabaRetailDeviceTradeShipAPIRequest struct {
	model.Params
	// 设备类型
	_deviceType string
	// 设备ID
	_deviceId string
	// 内部订单编号
	_tradeNo string
	// 详情
	_shipDetailList []ShipDetailDto
	// 选项
	_orderUpdateOption *OrderUpdateOption
}

// New
