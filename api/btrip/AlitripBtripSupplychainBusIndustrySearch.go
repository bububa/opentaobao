package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripsupplychainbusindustrysearch 汽车票行业搜索接口
// alitrip.btrip.supplychain.bus.industry.search
//
// 汽车票行业搜索接口
func Alitripbtripsupplychainbusindustrysearch(clt *core.SDKClient, req *btrip.AlitripbtripsupplychainbusindustrysearchAPIRequest, session string) (*btrip.AlitripbtripsupplychainbusindustrysearchAPIResponse, error) {
	var resp btrip.AlitripbtripsupplychainbusindustrysearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
