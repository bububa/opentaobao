package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TmallItemCombineGet
组合商品获取接口
tmall.item.combine.get

查询组合商品的SKU信息 */
func TmallItemCombineGet(clt *core.SDKClient, req *product.TmallItemCombineGetAPIRequest, session string) (*product.TmallItemCombineGetAPIResponse, error) {
	var resp product.TmallItemCombineGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
