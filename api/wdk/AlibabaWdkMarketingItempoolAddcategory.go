package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolAddcategory 增加商品池里面的类目
// alibaba.wdk.marketing.itempool.addcategory
//
// 增加商品池里面的类目
func AlibabaWdkMarketingItempoolAddcategory(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolAddcategoryAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolAddcategoryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
