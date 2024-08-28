package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantUserUpload 商家会员数据上传
// alibaba.tcls.aelophy.merchant.user.upload
//
// 商家会员数据上传
func AlibabaTclsAelophyMerchantUserUpload(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantUserUploadAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantUserUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
