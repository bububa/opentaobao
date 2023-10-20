package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Aliexpresssocialitemranking 社交排行榜
// aliexpress.social.item.ranking
//
// 社交商品成交排行榜
func Aliexpresssocialitemranking(clt *core.SDKClient, req *product.AliexpresssocialitemrankingAPIRequest, session string) (*product.AliexpresssocialitemrankingAPIResponse, error) {
	var resp product.AliexpresssocialitemrankingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
