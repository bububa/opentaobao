package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantOrderBatchUpload 商家订单数据批量上传
// alibaba.tcls.aelophy.merchant.order.batch.upload
//
// 商家订单数据上传
func AlibabaTclsAelophyMerchantOrderBatchUpload(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantOrderBatchUploadAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantOrderBatchUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
