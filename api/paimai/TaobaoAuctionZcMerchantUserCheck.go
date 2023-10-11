package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoAuctionZcMerchantUserCheck 通过手机号确认阿里资产商家
// taobao.auction.zc.merchant.user.check
//
// 通过手机号确认阿里资产商家
func TaobaoAuctionZcMerchantUserCheck(clt *core.SDKClient, req *paimai.TaobaoAuctionZcMerchantUserCheckAPIRequest, session string) (*paimai.TaobaoAuctionZcMerchantUserCheckAPIResponse, error) {
	var resp paimai.TaobaoAuctionZcMerchantUserCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
