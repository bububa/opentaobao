package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemseriesitemseriesinsertorupdate 商品系列增删改接口
// tmall.item.series.itemseries.insertorupdate
//
// 商品系列增删改接口
func Tmallitemseriesitemseriesinsertorupdate(clt *core.SDKClient, req *product.TmallitemseriesitemseriesinsertorupdateAPIRequest, session string) (*product.TmallitemseriesitemseriesinsertorupdateAPIResponse, error) {
	var resp product.TmallitemseriesitemseriesinsertorupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
