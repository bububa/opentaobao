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

// NewAlibabaItAlbumDeviceSendimageRequest 初始化AlibabaItAlbumDeviceSendimageAPIRequest对象
func NewAlibabaItAlbumDeviceSendimageRequest() *AlibabaItAlbumDeviceSendimageAPIRequest {
	return &AlibabaItAlbumDeviceSendimageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItAlbumDeviceSendimageAPIRequest) GetApiMethodName() string {
	return "alibaba.it.album.device.sendimage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItAlbumDeviceSendimageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Mac Setter
// 下发图片mac地址
func (r *AlibabaItAlbumDeviceSendimageAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// Get Mac Getter
func (r AlibabaItAlbumDeviceSendimageAPIRequest) GetMac() string {
	return r._mac
}
