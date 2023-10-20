package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantOrderUpload 商家订单数据上传
// alibaba.tcls.aelophy.merchant.order.upload
//
// 商家订单数据上传
func AlibabaTclsAelophyMerchantOrderUpload(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantOrderUploadAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantOrderUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
