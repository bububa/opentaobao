package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantUserUpload 商家会员数据上传
// alibaba.tcls.aelophy.merchant.user.upload
//
// 商家会员数据上传
func AlibabaTclsAelophyMerchantUserUpload(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantUserUploadAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantUserUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
