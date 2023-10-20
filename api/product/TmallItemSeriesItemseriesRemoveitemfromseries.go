package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemseriesitemseriesremoveitemfromseries 从商品系列中移除商品
// tmall.item.series.itemseries.removeitemfromseries
//
// 从商品系列中移除商品
func Tmallitemseriesitemseriesremoveitemfromseries(clt *core.SDKClient, req *product.TmallitemseriesitemseriesremoveitemfromseriesAPIRequest, session string) (*product.TmallitemseriesitemseriesremoveitemfromseriesAPIResponse, error) {
	var resp product.TmallitemseriesitemseriesremoveitemfromseriesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
