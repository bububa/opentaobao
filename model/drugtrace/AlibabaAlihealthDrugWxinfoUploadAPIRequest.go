package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugwxinfouploadAPIRequest 小程序数据回传 API请求
// alibaba.alihealth.drug.wxinfo.upload
//
// 小程序数据回传
type AlibabaalihealthdrugwxinfouploadAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo string
	// 店铺名称
	_shopInfo string
	// 售货员信息
	_salerInfo string
	// 渠道
	_isvChannel string
}

// NewAlibabaalihealthdrugwxinfouploadRequest 初始化AlibabaalihealthdrugwxinfouploadAPIRequest对象
func NewAlibabaalihealthdrugwxinfouploadRequest() *AlibabaalihealthdrugwxinfouploadAPIRequest {
	return &AlibabaalihealthdrugwxinfouploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugwxinfouploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.wxinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugwxinfouploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugwxinfouploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *AlibabaalihealthdrugwxinfouploadAPIRequest) SetUserInfo(_userInfo string) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r AlibabaalihealthdrugwxinfouploadAPIRequest) GetUserInfo() string {
	return r._userInfo
}

// SetShopInfo is ShopInfo Setter
// 店铺名称
func (r *AlibabaalihealthdrugwxinfouploadAPIRequest) SetShopInfo(_shopInfo string) error {
	r._shopInfo = _shopInfo
	r.Set("shop_info", _shopInfo)
	return nil
}

// GetShopInfo ShopInfo Getter
func (r AlibabaalihealthdrugwxinfouploadAPIRequest) GetShopInfo() string {
	return r._shopInfo
}

// SetSalerInfo is SalerInfo Setter
// 售货员信息
func (r *AlibabaalihealthdrugwxinfouploadAPIRequest) SetSalerInfo(_salerInfo string) error {
	r._salerInfo = _salerInfo
	r.Set("saler_info", _salerInfo)
	return nil
}

// GetSalerInfo SalerInfo Getter
func (r AlibabaalihealthdrugwxinfouploadAPIRequest) GetSalerInfo() string {
	return r._salerInfo
}

// SetIsvChannel is IsvChannel Setter
// 渠道
func (r *AlibabaalihealthdrugwxinfouploadAPIRequest) SetIsvChannel(_isvChannel string) error {
	r._isvChannel = _isvChannel
	r.Set("isv_channel", _isvChannel)
	return nil
}

// GetIsvChannel IsvChannel Getter
func (r AlibabaalihealthdrugwxinfouploadAPIRequest) GetIsvChannel() string {
	return r._isvChannel
}
