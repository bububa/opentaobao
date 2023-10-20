package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemMemberpriceUpdate 商品售价会员价修改
// alibaba.wdk.item.memberprice.update
//
// 商品售价会员价修改
func AlibabaWdkItemMemberpriceUpdate(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMemberpriceUpdateAPIRequest, resp *wdkitem.AlibabaWdkItemMemberpriceUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
