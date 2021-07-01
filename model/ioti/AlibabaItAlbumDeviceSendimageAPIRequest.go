package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItAlbumDeviceSendimageAPIRequest
相框设备厂测刷图接口 API请求
alibaba.it.album.device.sendimage

提供传入电子相框设备mac，mac需属于厂测白名单设备，将设备刷新为系统默认的厂测图片 */
type AlibabaItAlbumDeviceSendimageAPIRequest struct {
	model.Params
	// 下发图片mac地址
	_mac string
}

// New
