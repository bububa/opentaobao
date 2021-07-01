package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TaobaoItemUpdateListing
一口价商品上架
taobao.item.update.listing

* 单个商品上架<br/>* 输入的num_iid必须属于当前会话用户 */
func TaobaoItemUpdateListing(clt *core.SDKClient, req *product.TaobaoItemUpdateListingAPIRequest, session string) (*product.TaobaoItemUpdateListingAPIResponse, error) {
	var resp product.TaobaoItemUpdateListingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
