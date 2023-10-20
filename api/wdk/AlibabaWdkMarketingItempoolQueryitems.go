package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolQueryitems 查询商品池活动下的商品
// alibaba.wdk.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
func AlibabaWdkMarketingItempoolQueryitems(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolQueryitemsAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolQueryitemsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
