package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// AliexpressAffiliateProductSmartmatch 联盟物料智能推荐api
// aliexpress.affiliate.product.smartmatch
//
// 联盟物料算法智能推荐
func AliexpressAffiliateProductSmartmatch(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateProductSmartmatchAPIRequest, resp *aecreatives.AliexpressAffiliateProductSmartmatchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
