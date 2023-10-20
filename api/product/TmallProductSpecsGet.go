package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallproductspecsget 获取产品的规格信息
// tmall.product.specs.get
//
// 按product_id或品牌下载产品规格，返回一组的产品规格信息。
func Tmallproductspecsget(clt *core.SDKClient, req *product.TmallproductspecsgetAPIRequest, session string) (*product.TmallproductspecsgetAPIResponse, error) {
	var resp product.TmallproductspecsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
