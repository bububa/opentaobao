package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// AlibabaShopCategoryGet 指定店铺分类信息查询接口
// alibaba.shop.category.get
//
// 按照卖家身份查询指定分类信息
func AlibabaShopCategoryGet(clt *core.SDKClient, req *shop.AlibabaShopCategoryGetAPIRequest, resp *shop.AlibabaShopCategoryGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
