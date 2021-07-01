package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

/* AliexpressAffiliateImageSearch
图搜
aliexpress.affiliate.image.search

图片搜索接口 */
func AliexpressAffiliateImageSearch(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateImageSearchAPIRequest, session string) (*aecreatives.AliexpressAffiliateImageSearchAPIResponse, error) {
	var resp aecreatives.AliexpressAffiliateImageSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
