package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyMerchantUserUploadAPIRequest) Reset() {
	r._userInfoList = r._userInfoList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantUserUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.user.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantUserUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyMerchantUserUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaTclsAelophyMerchantUserUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyMerchantUserUploadRequest()
	},
}

// GetAlibabaTclsAelophyMerchantUserUploadRequest 从 sync.Pool 获取 AlibabaTclsAelophyMerchantUserUploadAPIRequest
func GetAlibabaTclsAelophyMerchantUserUploadAPIRequest() *AlibabaTclsAelophyMerchantUserUploadAPIRequest {
	return poolAlibabaTclsAelophyMerchantUserUploadAPIRequest.Get().(*AlibabaTclsAelophyMerchantUserUploadAPIRequest)
}

// ReleaseAlibabaTclsAelophyMerchantUserUploadAPIRequest 将 AlibabaTclsAelophyMerchantUserUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantUserUploadAPIRequest(v *AlibabaTclsAelophyMerchantUserUploadAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantUserUploadAPIRequest.Put(v)
}
