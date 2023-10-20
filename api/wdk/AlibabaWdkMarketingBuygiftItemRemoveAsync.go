package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingBuygiftItemRemoveAsync 批量删除买赠商品
// alibaba.wdk.marketing.buygift.item.remove.async
//
// 批量删除买赠商品
func AlibabaWdkMarketingBuygiftItemRemoveAsync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest, resp *wdk.AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
