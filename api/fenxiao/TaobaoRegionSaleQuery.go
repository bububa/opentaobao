package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoRegionSaleQuery 查询商品销售区域
// taobao.region.sale.query
//
// 查询商品销售区域
func TaobaoRegionSaleQuery(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoRegionSaleQueryAPIRequest, resp *fenxiao.TaobaoRegionSaleQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
