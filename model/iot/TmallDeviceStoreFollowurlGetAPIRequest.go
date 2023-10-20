package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalldevicestorefollowurlgetAPIRequest 获取店铺关注链接 API请求
// tmall.device.store.followurl.get
//
// 获取智能硬件上的关注店铺的URL
type TmalldevicestorefollowurlgetAPIRequest struct {
	model.Params
	// 设备DeviceCode
	_deviceCode string
	// 关注完成后的回调地址,需要是EWS地址
	_callbackUrl string
	// 页面banner的图片，如果没有传入，会使用系统默认图
	_bannerImg string
	// 是否同时关注天猫理想站
	_followRetailAccount bool
	// 是否使用长期链接
	_longterm bool
}

// NewTmalldevicestorefollowurlgetRequest 初始化TmalldevicestorefollowurlgetAPIRequest对象
func NewTmalldevicestorefollowurlgetRequest() *TmalldevicestorefollowurlgetAPIRequest {
	return &TmalldevicestorefollowurlgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalldevicestorefollowurlgetAPIRequest) GetApiMethodName() string {
	return "tmall.device.store.followurl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalldevicestorefollowurlgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalldevicestorefollowurlgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceCode is DeviceCode Setter
// 设备DeviceCode
func (r *TmalldevicestorefollowurlgetAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r TmalldevicestorefollowurlgetAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetCallbackUrl is CallbackUrl Setter
// 关注完成后的回调地址,需要是EWS地址
func (r *TmalldevicestorefollowurlgetAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// GetCallbackUrl CallbackUrl Getter
func (r TmalldevicestorefollowurlgetAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}

// SetBannerImg is BannerImg Setter
// 页面banner的图片，如果没有传入，会使用系统默认图
func (r *TmalldevicestorefollowurlgetAPIRequest) SetBannerImg(_bannerImg string) error {
	r._bannerImg = _bannerImg
	r.Set("banner_img", _bannerImg)
	return nil
}

// GetBannerImg BannerImg Getter
func (r TmalldevicestorefollowurlgetAPIRequest) GetBannerImg() string {
	return r._bannerImg
}

// SetFollowRetailAccount is FollowRetailAccount Setter
// 是否同时关注天猫理想站
func (r *TmalldevicestorefollowurlgetAPIRequest) SetFollowRetailAccount(_followRetailAccount bool) error {
	r._followRetailAccount = _followRetailAccount
	r.Set("follow_retail_account", _followRetailAccount)
	return nil
}

// GetFollowRetailAccount FollowRetailAccount Getter
func (r TmalldevicestorefollowurlgetAPIRequest) GetFollowRetailAccount() bool {
	return r._followRetailAccount
}

// SetLongterm is Longterm Setter
// 是否使用长期链接
func (r *TmalldevicestorefollowurlgetAPIRequest) SetLongterm(_longterm bool) error {
	r._longterm = _longterm
	r.Set("longterm", _longterm)
	return nil
}

// GetLongterm Longterm Getter
func (r TmalldevicestorefollowurlgetAPIRequest) GetLongterm() bool {
	return r._longterm
}
