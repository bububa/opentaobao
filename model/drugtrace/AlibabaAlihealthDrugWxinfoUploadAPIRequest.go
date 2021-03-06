package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugWxinfoUploadAPIRequest 小程序数据回传 API请求
// alibaba.alihealth.drug.wxinfo.upload
//
// 小程序数据回传
type AlibabaAlihealthDrugWxinfoUploadAPIRequest struct {
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

// NewAlibabaAlihealthDrugWxinfoUploadRequest 初始化AlibabaAlihealthDrugWxinfoUploadAPIRequest对象
func NewAlibabaAlihealthDrugWxinfoUploadRequest() *AlibabaAlihealthDrugWxinfoUploadAPIRequest {
	return &AlibabaAlihealthDrugWxinfoUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.wxinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *AlibabaAlihealthDrugWxinfoUploadAPIRequest) SetUserInfo(_userInfo string) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetUserInfo() string {
	return r._userInfo
}

// SetShopInfo is ShopInfo Setter
// 店铺名称
func (r *AlibabaAlihealthDrugWxinfoUploadAPIRequest) SetShopInfo(_shopInfo string) error {
	r._shopInfo = _shopInfo
	r.Set("shop_info", _shopInfo)
	return nil
}

// GetShopInfo ShopInfo Getter
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetShopInfo() string {
	return r._shopInfo
}

// SetSalerInfo is SalerInfo Setter
// 售货员信息
func (r *AlibabaAlihealthDrugWxinfoUploadAPIRequest) SetSalerInfo(_salerInfo string) error {
	r._salerInfo = _salerInfo
	r.Set("saler_info", _salerInfo)
	return nil
}

// GetSalerInfo SalerInfo Getter
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetSalerInfo() string {
	return r._salerInfo
}

// SetIsvChannel is IsvChannel Setter
// 渠道
func (r *AlibabaAlihealthDrugWxinfoUploadAPIRequest) SetIsvChannel(_isvChannel string) error {
	r._isvChannel = _isvChannel
	r.Set("isv_channel", _isvChannel)
	return nil
}

// GetIsvChannel IsvChannel Getter
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetIsvChannel() string {
	return r._isvChannel
}
