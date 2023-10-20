package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressPostproductRedefiningFindaeproductbyidfordropshipper Dropshipper查找商品信息接口
// aliexpress.postproduct.redefining.findaeproductbyidfordropshipper
//
// 提供给Dropshipper的通过商品ID查找商品信息的接口，只有特定买家可以使用
func AliexpressPostproductRedefiningFindaeproductbyidfordropshipper(clt *core.SDKClient, req *aedropshiper.AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest, resp *aedropshiper.AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
