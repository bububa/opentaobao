package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallDeviceShortchainurlGetAPIRequest
获取二维码短链接 API请求
tmall.device.shortchainurl.get

获取二维码短链接 */
type TmallDeviceShortchainurlGetAPIRequest struct {
	model.Params
	// 是否需要长期二维码，默认否
	_longterm bool
	// 需要生成短链接的url
	_url string
	// 设备DeviceCode
	_deviceCode string
	// 商户中心门店ID
	_storeId int64
	// 动作类型，支持自定义
	_action string
}

// New
