package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripsupplychainflightindustrysearch 机票行业搜索接口
// alitrip.btrip.supplychain.flight.industry.search
//
// 【商旅】机票行业搜索
func Alitripbtripsupplychainflightindustrysearch(clt *core.SDKClient, req *btrip.AlitripbtripsupplychainflightindustrysearchAPIRequest, session string) (*btrip.AlitripbtripsupplychainflightindustrysearchAPIResponse, error) {
	var resp btrip.AlitripbtripsupplychainflightindustrysearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
