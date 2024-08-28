package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoRegionPriceQuery 区域价格查询
// taobao.region.price.query
//
// 区域价格查询
func TaobaoRegionPriceQuery(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoRegionPriceQueryAPIRequest, resp *fenxiao.TaobaoRegionPriceQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
