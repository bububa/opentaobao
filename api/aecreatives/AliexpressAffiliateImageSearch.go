package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// AliexpressAffiliateImageSearch 图搜
// aliexpress.affiliate.image.search
//
// 图片搜索接口
func AliexpressAffiliateImageSearch(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateImageSearchAPIRequest, resp *aecreatives.AliexpressAffiliateImageSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
