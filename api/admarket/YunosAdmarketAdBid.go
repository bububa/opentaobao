package admarket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/admarket"
)

// Yunosadmarketadbid 广告竞价服务
// yunos.admarket.ad.bid
//
// 广告竞价服务
func Yunosadmarketadbid(clt *core.SDKClient, req *admarket.YunosadmarketadbidAPIRequest, session string) (*admarket.YunosadmarketadbidAPIResponse, error) {
	var resp admarket.YunosadmarketadbidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
