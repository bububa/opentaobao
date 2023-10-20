package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkTraceUrlGet 溯源url透出
// alibaba.wdk.trace.url.get
//
// 根据shopId和skuCode返回商品溯源url
func AlibabaWdkTraceUrlGet(clt *core.SDKClient, req *wdk.AlibabaWdkTraceUrlGetAPIRequest, resp *wdk.AlibabaWdkTraceUrlGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
