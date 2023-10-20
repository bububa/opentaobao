package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AliexpressSocialItemRanking 社交排行榜
// aliexpress.social.item.ranking
//
// 社交商品成交排行榜
func AliexpressSocialItemRanking(clt *core.SDKClient, req *product.AliexpressSocialItemRankingAPIRequest, resp *product.AliexpressSocialItemRankingAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
