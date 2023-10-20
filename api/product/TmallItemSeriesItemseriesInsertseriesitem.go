package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSeriesItemseriesInsertseriesitem 向系列中添加系列商品
// tmall.item.series.itemseries.insertseriesitem
//
// 向系列中添加系列商品
func TmallItemSeriesItemseriesInsertseriesitem(clt *core.SDKClient, req *product.TmallItemSeriesItemseriesInsertseriesitemAPIRequest, session string) (*product.TmallItemSeriesItemseriesInsertseriesitemAPIResponse, error) {
	var resp product.TmallItemSeriesItemseriesInsertseriesitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
