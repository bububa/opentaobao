package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemMerchantskuQuery 商家商品信息查询
// alibaba.wdk.item.merchantsku.query
//
// 商家商品信息查询
func AlibabaWdkItemMerchantskuQuery(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMerchantskuQueryAPIRequest, resp *wdkitem.AlibabaWdkItemMerchantskuQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
