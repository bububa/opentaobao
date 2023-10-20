package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolExcludeskucode 商品池排除商品【品类优惠使用】
// alibaba.wdk.marketing.itempool.excludeskucode
//
// 品类优惠新增排除池
func AlibabaWdkMarketingItempoolExcludeskucode(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
