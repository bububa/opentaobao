package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSeriesItemseriesRemoveitemfromseries 从商品系列中移除商品
// tmall.item.series.itemseries.removeitemfromseries
//
// 从商品系列中移除商品
func TmallItemSeriesItemseriesRemoveitemfromseries(clt *core.SDKClient, req *product.TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest, session string) (*product.TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse, error) {
	var resp product.TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
