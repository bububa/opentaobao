package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingBuygiftItemAddAsync 批量发布买赠商品
// alibaba.wdk.marketing.buygift.item.add.async
//
// 批量发布买赠商品
func AlibabaWdkMarketingBuygiftItemAddAsync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest, resp *wdk.AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
