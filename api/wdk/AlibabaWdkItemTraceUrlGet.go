package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkItemTraceUrlGet 根据shopId和skuCode返回商品静态溯源url
// alibaba.wdk.item.trace.url.get
//
// 根据shopId和skuCode返回商品静态溯源url
func AlibabaWdkItemTraceUrlGet(clt *core.SDKClient, req *wdk.AlibabaWdkItemTraceUrlGetAPIRequest, resp *wdk.AlibabaWdkItemTraceUrlGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
