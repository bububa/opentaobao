package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripsupplychaintrainindustrysearch 火车票行业搜索接口
// alitrip.btrip.supplychain.train.industry.search
//
// 【商旅】火车票行业搜索接口
func Alitripbtripsupplychaintrainindustrysearch(clt *core.SDKClient, req *btrip.AlitripbtripsupplychaintrainindustrysearchAPIRequest, session string) (*btrip.AlitripbtripsupplychaintrainindustrysearchAPIResponse, error) {
	var resp btrip.AlitripbtripsupplychaintrainindustrysearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
