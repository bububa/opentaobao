package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TmallSupplychainChannelProductRelease 供应商铺货
// tmall.supplychain.channel.product.release
//
// 供应商渠道铺货接口
func TmallSupplychainChannelProductRelease(ctx context.Context, clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductReleaseAPIRequest, resp *fenxiao.TmallSupplychainChannelProductReleaseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
