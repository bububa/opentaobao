package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// Aliexpressdsproductget 商品信息查询
// aliexpress.ds.product.get
//
// 商品信息查询
func Aliexpressdsproductget(clt *core.SDKClient, req *aedropshiper.AliexpressdsproductgetAPIRequest, session string) (*aedropshiper.AliexpressdsproductgetAPIResponse, error) {
	var resp aedropshiper.AliexpressdsproductgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
