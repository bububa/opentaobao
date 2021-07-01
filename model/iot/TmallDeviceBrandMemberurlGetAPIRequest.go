package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallDeviceBrandMemberurlGetAPIRequest
获取智能硬件旗舰店入会码 API请求
tmall.device.brand.memberurl.get

获取旗舰店在智能硬件上的入会码 */
type TmallDeviceBrandMemberurlGetAPIRequest struct {
	model.Params
	// 设备DeviceCode
	_deviceCode string
	// 入会后的回调地址
	_callbackUrl string
	// 是否使用长期链接
	_longterm bool
	// 页面banner的图片，如果没有传入，会使用系统默认图
	_bannerImg string
	// 是否同时关注天猫理想站
	_followRetailAccount bool
}

// New
