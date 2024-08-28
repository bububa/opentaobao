package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSeriesItemseriesInsertorupdate 商品系列增删改接口
// tmall.item.series.itemseries.insertorupdate
//
// 商品系列增删改接口
func TmallItemSeriesItemseriesInsertorupdate(ctx context.Context, clt *core.SDKClient, req *product.TmallItemSeriesItemseriesInsertorupdateAPIRequest, resp *product.TmallItemSeriesItemseriesInsertorupdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
