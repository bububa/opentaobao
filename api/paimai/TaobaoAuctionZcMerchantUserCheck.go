package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// Taobaoauctionzcmerchantusercheck 通过手机号确认阿里资产商家
// taobao.auction.zc.merchant.user.check
//
// 通过手机号确认阿里资产商家
func Taobaoauctionzcmerchantusercheck(clt *core.SDKClient, req *paimai.TaobaoauctionzcmerchantusercheckAPIRequest, session string) (*paimai.TaobaoauctionzcmerchantusercheckAPIResponse, error) {
	var resp paimai.TaobaoauctionzcmerchantusercheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
