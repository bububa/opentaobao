package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// AliexpressAffiliateProductdetailGet 联盟商品详情获取接口
// aliexpress.affiliate.productdetail.get
//
// 联盟推广商品搜索接口，用于搜索联盟推广商品数据
func AliexpressAffiliateProductdetailGet(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateProductdetailGetAPIRequest, resp *aecreatives.AliexpressAffiliateProductdetailGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
