package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Aliexpresssocialitempromotion 获取推广链接
// aliexpress.social.item.promotion
//
// 获取商品社交推广链接
func Aliexpresssocialitempromotion(clt *core.SDKClient, req *product.AliexpresssocialitempromotionAPIRequest, session string) (*product.AliexpresssocialitempromotionAPIResponse, error) {
	var resp product.AliexpresssocialitempromotionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
