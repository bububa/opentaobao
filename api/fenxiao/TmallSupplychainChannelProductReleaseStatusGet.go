package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TmallSupplychainChannelProductReleaseStatusGet 产品铺货状态查询
// tmall.supplychain.channel.product.release.status.get
//
// 巴拿马战役渠道产品状态查询
func TmallSupplychainChannelProductReleaseStatusGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductReleaseStatusGetAPIRequest, resp *fenxiao.TmallSupplychainChannelProductReleaseStatusGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
