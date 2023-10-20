package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingOpenVersionCount 版本数量查询
// alibaba.wdk.marketing.open.version.count
//
// 版本数量查询
func AlibabaWdkMarketingOpenVersionCount(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenVersionCountAPIRequest, resp *wdk.AlibabaWdkMarketingOpenVersionCountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
