package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsProductGet 商品信息查询
// aliexpress.ds.product.get
//
// 商品信息查询
func AliexpressDsProductGet(clt *core.SDKClient, req *aedropshiper.AliexpressDsProductGetAPIRequest, resp *aedropshiper.AliexpressDsProductGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
