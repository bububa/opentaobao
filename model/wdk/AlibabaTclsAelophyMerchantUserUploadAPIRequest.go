package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyMerchantUserUploadAPIRequest
商家会员数据上传 API请求
alibaba.tcls.aelophy.merchant.user.upload

商家会员数据上传 */
type AlibabaTclsAelophyMerchantUserUploadAPIRequest struct {
	model.Params
	// 渠道用户信息
	_userInfoList []MerchantUserInfo
}

// New
