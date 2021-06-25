package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家会员数据上传 APIRequest
alibaba.tcls.aelophy.merchant.user.upload

商家会员数据上传
*/
type AlibabaTclsAelophyMerchantUserUploadRequest struct {
    model.Params

    // 渠道用户信息
    userInfoList   []MerchantUserInfo 

}

func NewAlibabaTclsAelophyMerchantUserUploadRequest() *AlibabaTclsAelophyMerchantUserUploadRequest{
    return &AlibabaTclsAelophyMerchantUserUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyMerchantUserUploadRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.user.upload"
}

func (r AlibabaTclsAelophyMerchantUserUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyMerchantUserUploadRequest) SetUserInfoList(userInfoList []MerchantUserInfo) error {
    r.userInfoList = userInfoList
    r.Set("user_info_list", userInfoList)
    return nil
}

func (r AlibabaTclsAelophyMerchantUserUploadRequest) GetUserInfoList() []MerchantUserInfo {
    return r.userInfoList
}

