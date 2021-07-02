package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantUserUploadAPIRequest 商家会员数据上传 API请求
// alibaba.tcls.aelophy.merchant.user.upload
//
// 商家会员数据上传
type AlibabaTclsAelophyMerchantUserUploadAPIRequest struct {
	model.Params
	// 渠道用户信息
	_userInfoList []MerchantUserInfo
}

// NewAlibabaTclsAelophyMerchantUserUploadRequest 初始化AlibabaTclsAelophyMerchantUserUploadAPIRequest对象
func NewAlibabaTclsAelophyMerchantUserUploadRequest() *AlibabaTclsAelophyMerchantUserUploadAPIRequest {
	return &AlibabaTclsAelophyMerchantUserUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantUserUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.user.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantUserUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserInfoList is UserInfoList Setter
// 渠道用户信息
func (r *AlibabaTclsAelophyMerchantUserUploadAPIRequest) SetUserInfoList(_userInfoList []MerchantUserInfo) error {
	r._userInfoList = _userInfoList
	r.Set("user_info_list", _userInfoList)
	return nil
}

// GetUserInfoList UserInfoList Getter
func (r AlibabaTclsAelophyMerchantUserUploadAPIRequest) GetUserInfoList() []MerchantUserInfo {
	return r._userInfoList
}
