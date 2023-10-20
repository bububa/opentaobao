package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripsupplychaintraincity 火车站数据查询
// alitrip.btrip.supplychain.train.city
//
// 火车站数据查询
func Alitripbtripsupplychaintraincity(clt *core.SDKClient, req *btrip.AlitripbtripsupplychaintraincityAPIRequest, session string) (*btrip.AlitripbtripsupplychaintraincityAPIResponse, error) {
	var resp btrip.AlitripbtripsupplychaintraincityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
