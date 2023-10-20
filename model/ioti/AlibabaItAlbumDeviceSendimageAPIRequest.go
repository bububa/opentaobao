package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitalbumdevicesendimageAPIRequest 相框设备厂测刷图接口 API请求
// alibaba.it.album.device.sendimage
//
// 提供传入电子相框设备mac，mac需属于厂测白名单设备，将设备刷新为系统默认的厂测图片
type AlibabaitalbumdevicesendimageAPIRequest struct {
	model.Params
	// 下发图片mac地址
	_mac string
}

// NewAlibabaitalbumdevicesendimageRequest 初始化AlibabaitalbumdevicesendimageAPIRequest对象
func NewAlibabaitalbumdevicesendimageRequest() *AlibabaitalbumdevicesendimageAPIRequest {
	return &AlibabaitalbumdevicesendimageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaitalbumdevicesendimageAPIRequest) GetApiMethodName() string {
	return "alibaba.it.album.device.sendimage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaitalbumdevicesendimageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaitalbumdevicesendimageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMac is Mac Setter
// 下发图片mac地址
func (r *AlibabaitalbumdevicesendimageAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaitalbumdevicesendimageAPIRequest) GetMac() string {
	return r._mac
}
