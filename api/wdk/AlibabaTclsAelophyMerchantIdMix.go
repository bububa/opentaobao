package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantIdMix 商家用户id混淆
// alibaba.tcls.aelophy.merchant.id.mix
//
// 商家用户id混淆
func AlibabaTclsAelophyMerchantIdMix(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantIdMixAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantIdMixAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
