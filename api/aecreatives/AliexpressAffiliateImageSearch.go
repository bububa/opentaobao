package aecreatives

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// AliexpressAffiliateImageSearch 图搜
// aliexpress.affiliate.image.search
//
// 图片搜索接口
func AliexpressAffiliateImageSearch(ctx context.Context, clt *core.SDKClient, req *aecreatives.AliexpressAffiliateImageSearchAPIRequest, resp *aecreatives.AliexpressAffiliateImageSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
