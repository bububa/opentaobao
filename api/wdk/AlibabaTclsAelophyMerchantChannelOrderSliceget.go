package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantChannelOrderSliceget 获取运力时间片信息
// alibaba.tcls.aelophy.merchant.channel.order.sliceget
//
// 获取履约时间片
func AlibabaTclsAelophyMerchantChannelOrderSliceget(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
