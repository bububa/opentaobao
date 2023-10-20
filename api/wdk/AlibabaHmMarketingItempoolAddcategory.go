package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolAddcategory 增加商品池里面的类目
// alibaba.hm.marketing.itempool.addcategory
//
// 增加商品池里面的类目
func AlibabaHmMarketingItempoolAddcategory(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolAddcategoryAPIRequest, resp *wdk.AlibabaHmMarketingItempoolAddcategoryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
