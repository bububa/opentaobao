package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallDeviceStoreFollowurlGetAPIRequest
获取店铺关注链接 API请求
tmall.device.store.followurl.get

获取智能硬件上的关注店铺的URL */
type TmallDeviceStoreFollowurlGetAPIRequest struct {
	model.Params
	// 设备DeviceCode
	_deviceCode string
	// 关注完成后的回调地址,需要是EWS地址
	_callbackUrl string
	// 是否同时关注天猫理想站
	_followRetailAccount bool
	// 是否使用长期链接
	_longterm bool
	// 页面banner的图片，如果没有传入，会使用系统默认图
	_bannerImg string
}

// New
