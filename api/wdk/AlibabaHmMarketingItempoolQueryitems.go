package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolQueryitems 查询商品池活动下面的商品
// alibaba.hm.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
func AlibabaHmMarketingItempoolQueryitems(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolQueryitemsAPIRequest, resp *wdk.AlibabaHmMarketingItempoolQueryitemsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
