package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Aliexpresssocialitemsearch AE社交选品
// aliexpress.social.item.search
//
// AE社交选品,通过各种筛选条件对社交商品池进行筛选
func Aliexpresssocialitemsearch(clt *core.SDKClient, req *product.AliexpresssocialitemsearchAPIRequest, session string) (*product.AliexpresssocialitemsearchAPIResponse, error) {
	var resp product.AliexpresssocialitemsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
