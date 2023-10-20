package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AliexpressSocialItemSearch AE社交选品
// aliexpress.social.item.search
//
// AE社交选品,通过各种筛选条件对社交商品池进行筛选
func AliexpressSocialItemSearch(clt *core.SDKClient, req *product.AliexpressSocialItemSearchAPIRequest, session string) (*product.AliexpressSocialItemSearchAPIResponse, error) {
	var resp product.AliexpressSocialItemSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
