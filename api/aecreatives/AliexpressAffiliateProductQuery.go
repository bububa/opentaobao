package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// AliexpressAffiliateProductQuery 联盟推广商品获取接口
// aliexpress.affiliate.product.query
//
// 联盟推广商品搜索接口，用于搜索联盟推广商品数据
func AliexpressAffiliateProductQuery(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateProductQueryAPIRequest, session string) (*aecreatives.AliexpressAffiliateProductQueryAPIResponse, error) {
	var resp aecreatives.AliexpressAffiliateProductQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
