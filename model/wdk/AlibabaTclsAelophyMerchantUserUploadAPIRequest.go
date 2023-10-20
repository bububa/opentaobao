package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantuseruploadAPIRequest 商家会员数据上传 API请求
// alibaba.tcls.aelophy.merchant.user.upload
//
// 商家会员数据上传
type AlibabatclsaelophymerchantuseruploadAPIRequest struct {
	model.Params
	// 渠道用户信息
	_userInfoList []MerchantUserInfo
}

// NewAlibabatclsaelophymerchantuseruploadRequest 初始化AlibabatclsaelophymerchantuseruploadAPIRequest对象
func NewAlibabatclsaelophymerchantuseruploadRequest() *AlibabatclsaelophymerchantuseruploadAPIRequest {
	return &AlibabatclsaelophymerchantuseruploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophymerchantuseruploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.user.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophymerchantuseruploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophymerchantuseruploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserInfoList is UserInfoList Setter
// 渠道用户信息
func (r *AlibabatclsaelophymerchantuseruploadAPIRequest) SetUserInfoList(_userInfoList []MerchantUserInfo) error {
	r._userInfoList = _userInfoList
	r.Set("user_info_list", _userInfoList)
	return nil
}

// GetUserInfoList UserInfoList Getter
func (r AlibabatclsaelophymerchantuseruploadAPIRequest) GetUserInfoList() []MerchantUserInfo {
	return r._userInfoList
}
