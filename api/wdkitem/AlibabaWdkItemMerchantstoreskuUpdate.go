package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemMerchantstoreskuUpdate 门店商品信息修改
// alibaba.wdk.item.merchantstoresku.update
//
// 门店商品信息修改
func AlibabaWdkItemMerchantstoreskuUpdate(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMerchantstoreskuUpdateAPIRequest, resp *wdkitem.AlibabaWdkItemMerchantstoreskuUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
