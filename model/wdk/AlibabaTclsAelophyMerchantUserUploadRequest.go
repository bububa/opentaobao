package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家会员数据上传 API请求
alibaba.tcls.aelophy.merchant.user.upload

商家会员数据上传
*/
type AlibabaTclsAelophyMerchantUserUploadRequest struct {
    model.Params
    // 渠道用户信息
    userInfoList   []MerchantUserInfo
}

// 初始化AlibabaTclsAelophyMerchantUserUploadRequest对象
func NewAlibabaTclsAelophyMerchantUserUploadRequest() *AlibabaTclsAelophyMerchantUserUploadRequest{
    return &AlibabaTclsAelophyMerchantUserUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantUserUploadRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.user.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantUserUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserInfoList Setter
// 渠道用户信息
func (r *AlibabaTclsAelophyMerchantUserUploadRequest) SetUserInfoList(userInfoList []MerchantUserInfo) error {
    r.userInfoList = userInfoList
    r.Set("user_info_list", userInfoList)
    return nil
}

// UserInfoList Getter
func (r AlibabaTclsAelophyMerchantUserUploadRequest) GetUserInfoList() []MerchantUserInfo {
    return r.userInfoList
}
