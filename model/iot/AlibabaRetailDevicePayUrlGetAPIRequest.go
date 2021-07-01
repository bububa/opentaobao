package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailDevicePayUrlGetAPIRequest
贩卖机支付二维链接获取 API请求
alibaba.retail.device.payUrl.get

贩卖机支付二维链接获取 */
type AlibabaRetailDevicePayUrlGetAPIRequest struct {
	model.Params
	// 外部订单id
	_isvOrderId string
	// 业务名称
	_bizName string
	// 商品id
	_itemId int64
	// 设备sn
	_deviceId string
	// 1表示商品box，0或者为空表示普通商品
	_itemType int64
}

// New
