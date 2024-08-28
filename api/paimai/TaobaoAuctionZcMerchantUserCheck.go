package paimai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoAuctionZcMerchantUserCheck 通过手机号确认阿里资产商家
// taobao.auction.zc.merchant.user.check
//
// 通过手机号确认阿里资产商家
func TaobaoAuctionZcMerchantUserCheck(ctx context.Context, clt *core.SDKClient, req *paimai.TaobaoAuctionZcMerchantUserCheckAPIRequest, resp *paimai.TaobaoAuctionZcMerchantUserCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
