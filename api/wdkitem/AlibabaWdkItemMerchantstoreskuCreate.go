package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemMerchantstoreskuCreate 门店商品信息新建
// alibaba.wdk.item.merchantstoresku.create
//
// 门店商品信息新建
func AlibabaWdkItemMerchantstoreskuCreate(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMerchantstoreskuCreateAPIRequest, resp *wdkitem.AlibabaWdkItemMerchantstoreskuCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
