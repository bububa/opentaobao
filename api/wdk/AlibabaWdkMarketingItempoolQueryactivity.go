package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolQueryactivity 查找商品池活动
// alibaba.wdk.marketing.itempool.queryactivity
//
// 查找商品池活动
func AlibabaWdkMarketingItempoolQueryactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolQueryactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolQueryactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
