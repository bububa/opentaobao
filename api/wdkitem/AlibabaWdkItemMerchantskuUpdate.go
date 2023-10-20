package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemMerchantskuUpdate 商家商品修改
// alibaba.wdk.item.merchantsku.update
//
// 商家商品修改
func AlibabaWdkItemMerchantskuUpdate(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMerchantskuUpdateAPIRequest, resp *wdkitem.AlibabaWdkItemMerchantskuUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
