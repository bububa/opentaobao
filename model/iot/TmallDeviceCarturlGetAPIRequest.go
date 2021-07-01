package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallDeviceCarturlGetAPIRequest
添加商品到购物车 API请求
tmall.device.carturl.get

获取二维码，支持添加商品到购物车 */
type TmallDeviceCarturlGetAPIRequest struct {
	model.Params
	// 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
	_itemIds []string
	// 设备业务编码
	_deviceCode string
	// 是否申请长期链接
	_longterm bool
}

// New
