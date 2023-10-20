package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSeriesItemseriesRemoveitemfromseries 从商品系列中移除商品
// tmall.item.series.itemseries.removeitemfromseries
//
// 从商品系列中移除商品
func TmallItemSeriesItemseriesRemoveitemfromseries(clt *core.SDKClient, req *product.TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest, resp *product.TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
