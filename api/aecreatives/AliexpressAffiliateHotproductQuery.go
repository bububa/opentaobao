package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// AliexpressAffiliateHotproductQuery 查询联盟爆品数据
// aliexpress.affiliate.hotproduct.query
//
// 查询联盟爆品API
func AliexpressAffiliateHotproductQuery(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateHotproductQueryAPIRequest, resp *aecreatives.AliexpressAffiliateHotproductQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
