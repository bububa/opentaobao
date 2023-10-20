package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemseriesitemseriesinsertseriesitem 向系列中添加系列商品
// tmall.item.series.itemseries.insertseriesitem
//
// 向系列中添加系列商品
func Tmallitemseriesitemseriesinsertseriesitem(clt *core.SDKClient, req *product.TmallitemseriesitemseriesinsertseriesitemAPIRequest, session string) (*product.TmallitemseriesitemseriesinsertseriesitemAPIResponse, error) {
	var resp product.TmallitemseriesitemseriesinsertseriesitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
