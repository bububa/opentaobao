package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtPayMerchantStallSigningModify 三级商户进件修改
// tmall.nrt.pay.merchant.stall.signing.modify
//
// 三级商户进件修改
func TmallNrtPayMerchantStallSigningModify(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtPayMerchantStallSigningModifyAPIRequest, resp *nrt.TmallNrtPayMerchantStallSigningModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
